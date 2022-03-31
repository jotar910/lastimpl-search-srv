package store

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"lastimplementation.com/pkg/services/projects"
	"lastimplementation.com/pkg/services/projects/logger"
	"lastimplementation.com/pkg/services/projects/models"
	"lastimplementation.com/pkg/services/projects/store/dao"
	"lastimplementation.com/pkg/services/projects/store/queries"
)

type projectsRepo struct {
	l  logger.Logger
	db *sql.DB
}

// New creates a projects repo.
func New(l logger.Logger, db *sql.DB) *projectsRepo {
	return &projectsRepo{l.WithPrefix("store"), db}
}

// Reset resets the projects tables.
func (pr *projectsRepo) Reset(ctx context.Context) error {
	log := pr.l.WithPrefix("reset")

	log.Trace("Resetting the projects repo")
	if _, err := pr.db.Exec(queries.QueryDropTables); err != nil {
		return fmt.Errorf("executing query: %w", err)
	}
	if _, err := pr.db.Exec(queries.QueryCreateTables); err != nil {
		return fmt.Errorf("executing query: %w", err)
	}
	return nil
}

// Add inserts a new project.
func (pr *projectsRepo) Add(ctx context.Context, project models.Project) (int, error) {
	log := pr.l.WithPrefix("add")

	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		log.Error("begining transaction", err.Error())
		return -1, err
	}

	p := dao.Project{Name: project.Name, Description: project.Description}
	if err = p.Insert(ctx, tx, boil.Infer()); err != nil {
		log.Error("inserting new project", err)
		tx.Rollback()
		if strings.HasSuffix(err.Error(), ErrDuplicated("projects_name_key")) {
			return -1, projects.ErrAddProjectDuplicatedName
		}
		return -1, err
	}

	if err := pr.addTags(ctx, tx, p.ID, project.Tags); err != nil {
		log.Error("adding tags to the project", err.Error())
		tx.Rollback()
		return -1, err
	}

	if err := pr.addFiles(ctx, tx, p.ID, project.Files); err != nil {
		log.Error("adding code files to the project", err.Error())
		tx.Rollback()
		return -1, err
	}

	tx.Commit()
	return p.ID, nil
}

func (pr *projectsRepo) addFiles(ctx context.Context, tx *sql.Tx, pId int, cfs []models.CodeFile) error {
	if len(cfs) == 0 {
		return nil
	}
	for _, cf := range cfs {
		file := dao.CodeFile{
			ProjectID: pId,
			Name:      cf.Name,
			Content:   cf.Content,
		}
		if err := file.Insert(context.Background(), tx, boil.Infer()); err != nil {
			return fmt.Errorf("inserting project %q: %w", cf.Name, err)
		}
	}
	return nil
}

func (pr *projectsRepo) addTags(ctx context.Context, tx *sql.Tx, pId int, tags []models.TagType) error {
	if len(tags) == 0 {
		return nil
	}
	existingTags, tagsToAdd, err := pr.getTagsToAdd(context.Background(), tx, tags)
	if err != nil {
		return err
	}
	for _, tag := range tagsToAdd {
		t := dao.Tag{Name: string(tag)}
		if err := t.Insert(context.Background(), tx, boil.Infer()); err != nil {
			return err
		}
		existingTags = append(existingTags, &t)
	}
	for _, tag := range existingTags {
		pt := dao.ProjectsTag{
			ProjectID: pId,
			TagID:     tag.ID,
		}
		if err := pt.Insert(context.Background(), tx, boil.Infer()); err != nil {
			return err
		}
	}
	return nil
}

func (pr *projectsRepo) getTagsToAdd(ctx context.Context, tx *sql.Tx, tags []models.TagType) ([]*dao.Tag, []models.TagType, error) {
	tagsStr := make([]interface{}, len(tags))
	for i, tag := range tags {
		tagsStr[i] = tag
	}
	pTags, err := dao.Tags(qm.WhereIn("name IN ?", tagsStr...)).All(ctx, tx)
	if err != nil {
		return nil, nil, err
	}
	pTagsMap := make(map[models.TagType]struct{})
	existingTags := make([]*dao.Tag, 0)
	for _, pTag := range pTags {
		pTagsMap[models.TagType(pTag.Name)] = struct{}{}
		existingTags = append(existingTags, pTag)
	}
	var leftTags []models.TagType
	for _, tag := range tags {
		if _, ok := pTagsMap[tag]; !ok {
			leftTags = append(leftTags, tag)
			pTagsMap[tag] = struct{}{}
		}
	}
	return existingTags, leftTags, nil
}

// Get fetches a specific project.
func (pr *projectsRepo) Get(ctx context.Context, id int) (models.Project, error) {
	log := pr.l.WithPrefix("get")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		log.Error("begining transaction", err)
		return models.Project{}, err
	}

	p, err := dao.Projects(
		qm.Where("id = ?", id),
		qm.Select(dao.ProjectColumns.ID, dao.ProjectColumns.Name, dao.ProjectColumns.Description),
		qm.Load(dao.ProjectRels.CodeFiles,
			qm.Select(dao.CodeFileColumns.ID, dao.CodeFileColumns.Name, dao.CodeFileColumns.Content, dao.CodeFileColumns.CreatedAt),
			qm.OrderBy(dao.CodeFileColumns.CreatedAt)),
		qm.Load(qm.Rels(dao.ProjectRels.ProjectsTags), qm.Select(dao.ProjectsTagColumns.TagID)),
		qm.Load(qm.Rels(dao.ProjectRels.ProjectsTags, dao.ProjectsTagRels.Tag),
			qm.Select(dao.TagColumns.ID, dao.TagColumns.Name)),
	).One(ctx, tx)
	if err != nil {
		log.Error("executing query to get project:", err)
		tx.Rollback()
		if strings.HasSuffix(err.Error(), ErrNotResult()) {
			return models.Project{}, projects.ErrProjectNotFound
		}
		return models.Project{}, err
	}
	tx.Commit()

	res := models.Project{
		Id: p.ID,
		ProjectDetails: models.ProjectDetails{
			Name:        p.Name,
			Description: p.Description,
		},
	}

	for _, tag := range p.R.ProjectsTags {
		res.Tags = append(res.Tags, models.TagType(tag.R.Tag.Name))
	}

	for _, cf := range p.R.CodeFiles {
		res.Files = append(res.Files, models.CodeFile{
			Id:      cf.ID,
			Name:    cf.Name,
			Content: cf.Content,
		})
	}

	return res, nil
}

// GetAll fetches the projects list.
func (pr *projectsRepo) GetAll(ctx context.Context, query string, page, limit int) (models.ProjectsList, error) {
	log := pr.l.WithPrefix("getAll")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		log.Error("begining transaction", err)
		return models.ProjectsList{}, err
	}

	projects, err := dao.Projects(
		qm.Select(dao.ProjectColumns.ID, dao.ProjectColumns.Name, dao.ProjectColumns.Description, dao.ProjectColumns.UpdatedAt),
		qm.Where("LOWER(name) ~ ?", strings.ToLower(query)),
		qm.Offset((page-1)*limit),
		qm.Limit(limit),
		qm.OrderBy(dao.ProjectColumns.Name),
		qm.Load(dao.ProjectRels.CodeFiles,
			qm.Select(dao.CodeFileColumns.ProjectID, dao.CodeFileColumns.ID, dao.CodeFileColumns.Name, dao.CodeFileColumns.CreatedAt),
			qm.OrderBy(dao.CodeFileColumns.CreatedAt)),
		qm.Load(dao.ProjectRels.ProjectsTags,
			qm.Select(dao.ProjectsTagColumns.ProjectID, dao.ProjectsTagColumns.TagID)),
		qm.Load(qm.Rels(dao.ProjectRels.ProjectsTags, dao.ProjectsTagRels.Tag),
			qm.Select(dao.TagColumns.ID, dao.TagColumns.Name)),
	).All(ctx, tx)
	if err != nil {
		log.Error("getting project items", err)
		tx.Rollback()
		return models.ProjectsList{}, err
	}

	count, err := dao.Projects().Count(ctx, pr.db)
	if err != nil {
		log.Error("counting total project items", err)
		tx.Rollback()
		return models.ProjectsList{}, err
	}

	tx.Commit()

	projectList := models.ProjectsList{Data: make([]models.ProjectItem, len(projects))}

	for i, p := range projects {
		projectList.Data[i] = models.ProjectItem{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			UpdatedAt:   p.UpdatedAt.Local().Unix(),
			Files:       make([]models.ProjectItemFile, len(p.R.CodeFiles)),
			Tags:        make([]models.Tag, len(p.R.ProjectsTags)),
		}
		for j, cf := range p.R.CodeFiles {
			projectList.Data[i].Files[j] = models.ProjectItemFile{
				Id:   cf.ID,
				Name: cf.Name,
			}
		}
		for j, tag := range p.R.ProjectsTags {
			projectList.Data[i].Tags[j] = models.Tag{
				Id:   tag.TagID,
				Name: models.TagType(tag.R.Tag.Name),
			}
		}
	}

	projectList.TotalItems = int(count)
	projectList.Page = page
	projectList.Count = len(projectList.Data)
	projectList.TotalPages = int(math.Ceil(float64(projectList.TotalItems) / float64(limit)))

	return projectList, nil
}

// Update updates the data from an existing project.
func (pr *projectsRepo) Update(ctx context.Context, id int, project models.ProjectDetails) error {
	log := pr.l.WithPrefix("update")

	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		log.Error("begining transaction", err)
		return err
	}

	columns := []string{dao.ProjectColumns.UpdatedAt}
	p := dao.Project{ID: id}
	if project.Name != "" {
		p.Name = project.Name
		columns = append(columns, dao.ProjectColumns.Name)
	}
	if project.Description != "" {
		p.Description = project.Description
		columns = append(columns, dao.ProjectColumns.Description)
	}
	changed, err := p.Update(ctx, tx, boil.Whitelist(columns...))
	if err != nil {
		log.Error("updating existing project", err)
		tx.Rollback()
		return err
	}
	if changed == 0 {
		err := projects.ErrProjectNotFound
		log.Error("updating existing project", err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// Delete deletes an existing project.
func (pr *projectsRepo) Delete(ctx context.Context, id int) error {
	log := pr.l.WithPrefix("delete")

	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		log.Error("begining transaction", err)
		return err
	}

	p, err := dao.Projects(
		qm.Select(dao.ProjectColumns.ID),
		qm.Where("id = ?", id),
	).One(ctx, tx)
	if err != nil {
		log.Error("finding project", err)
		tx.Rollback()
		if strings.HasSuffix(err.Error(), ErrNotResult()) {
			return projects.ErrProjectNotFound
		}
		return err
	}

	if _, err := dao.CodeFiles(qm.Where("project_id = ?", id)).DeleteAll(ctx, tx); err != nil {
		log.Error("deleting code files from existing project", err)
		tx.Rollback()
		return err
	}

	if _, err := dao.ProjectsTags(qm.Where("project_id = ?", id)).DeleteAll(ctx, tx); err != nil {
		log.Error("deleting code files from existing project", err)
		tx.Rollback()
		return err
	}

	dbProjectsHistory, err := dao.ProjectsHistories(
		qm.Select(dao.ProjectsHistoryColumns.ID, dao.ProjectsHistoryColumns.ProjectID),
		qm.Where("project_id = ?", id),
	).All(ctx, tx)
	if err != nil {
		log.Error("searching for the history of the project to delete")
		tx.Rollback()
		return err
	}

	dbProjectsHistoryId := make([]interface{}, len(dbProjectsHistory))
	for i, dbProjectHistory := range dbProjectsHistory {
		dbProjectsHistoryId[i] = dbProjectHistory.ID
	}

	if _, err := dao.ProjectsCodeFilesHistories(
		qm.Select(dao.ProjectsCodeFilesHistoryColumns.ID, dao.ProjectsCodeFilesHistoryColumns.RevisionID),
		qm.WhereIn("revision_id IN ?", dbProjectsHistoryId...),
	).DeleteAll(ctx, tx); err != nil {
		log.Error("deleting code files history for the deleted project", err)
		tx.Rollback()
		return err
	}

	if _, err := dbProjectsHistory.DeleteAll(ctx, tx); err != nil {
		log.Error("deleting project revision history for the deleted project", err)
		tx.Rollback()
		return err
	}

	if _, err := p.Delete(ctx, tx); err != nil {
		log.Error("deleting existing project", err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// GetFiles fetches all the files for a given project.
func (pr *projectsRepo) GetFiles(ctx context.Context, projectId int) (models.CodeFiles, error) {
	log := pr.l.WithPrefix("getFiles")

	dbFiles, err := dao.CodeFiles(
		qm.Select(dao.CodeFileColumns.ID, dao.CodeFileColumns.ProjectID, dao.CodeFileColumns.Name, dao.CodeFileColumns.Content, dao.CodeFileColumns.CreatedAt),
		qm.Where("project_id = ?", projectId),
		qm.Limit(models.MaximumCodeFiles),
		qm.OrderBy(dao.CodeFileColumns.CreatedAt),
	).All(ctx, pr.db)
	if err != nil {
		log.Error("fetching code file for an existing project", err)
		return nil, err
	}
	files := make([]models.CodeFile, len(dbFiles))
	for i, dbFile := range dbFiles {
		files[i] = models.CodeFile{
			Id:      dbFile.ID,
			Name:    dbFile.Name,
			Content: dbFile.Content,
		}
	}
	return files, nil
}

// UpdateFiles updates the files data for a given project.
func (pr *projectsRepo) UpdateFiles(ctx context.Context, projectId int, files []models.CodeFile) error {
	log := pr.l.WithPrefix("updateFiles")

	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		log.Error("begining transaction", err)
		return err
	}

	_, err = dao.Projects(
		qm.Select(dao.ProjectColumns.ID),
		qm.Where("id = ?", projectId),
	).One(ctx, tx)
	if err != nil {
		log.Error("finding project", err)
		tx.Rollback()
		if strings.HasSuffix(err.Error(), ErrNotResult()) {
			return projects.ErrProjectNotFound
		}
		return err
	}

	dbHistProj := dao.ProjectsHistory{ProjectID: projectId}
	if err := dbHistProj.Insert(ctx, tx, boil.Infer()); err != nil {
		log.Error("inserting project revision history", err)
		tx.Rollback()
		return err
	}

	dbFiles, err := dao.CodeFiles(qm.Where("project_id = ?", projectId)).All(ctx, tx)
	if err != nil {
		log.Error("getting project files")
		tx.Rollback()
		return err
	}

	for _, dbFile := range dbFiles {
		dbHistFile := dao.ProjectsCodeFilesHistory{
			Name:       dbFile.Name,
			Content:    dbFile.Content,
			RevisionID: dbHistProj.ID,
		}
		if err := dbHistFile.Insert(ctx, tx, boil.Infer()); err != nil {
			log.Error("inserting code file to history", err)
			tx.Rollback()
			return err
		}
	}

	if err := pr.mergeFiles(ctx, tx, projectId, files, dbFiles); err != nil {
		log.Error("merging project files", err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (pr *projectsRepo) mergeFiles(ctx context.Context, tx *sql.Tx, projectId int, files []models.CodeFile, dbFiles dao.CodeFileSlice) error {
	filesToUpdateMap := make(map[int]models.CodeFile)
	var filesToAdd []models.CodeFile
	for _, file := range files {
		if file.Id != 0 {
			filesToUpdateMap[file.Id] = file
		} else {
			filesToAdd = append(filesToAdd, file)
		}
	}

	if len(filesToUpdateMap) == 0 {
		return pr.setFiles(ctx, tx, projectId, files, dbFiles)
	}

	var dbFilesToDelete dao.CodeFileSlice
	var dbFilesToUpdate dao.CodeFileSlice
	for _, dbFile := range dbFiles {
		if _, ok := filesToUpdateMap[dbFile.ID]; ok {
			dbFilesToUpdate = append(dbFilesToUpdate, dbFile)
		} else {
			dbFilesToDelete = append(dbFilesToDelete, dbFile)
		}
	}

	if _, err := dbFilesToDelete.DeleteAll(ctx, tx); err != nil {
		return fmt.Errorf("deleting some existing project files: %w", err)
	}

	for _, dbFile := range dbFilesToUpdate {
		file := filesToUpdateMap[dbFile.ID]
		dbFile.Name = file.Name
		dbFile.Content = file.Content
		if _, err := dbFile.Update(ctx, tx, boil.Whitelist(dao.CodeFileColumns.Name, dao.CodeFileColumns.Content)); err != nil {
			return fmt.Errorf("updating existing project file %d: %w", dbFile.ID, err)
		}
	}

	if err := pr.addFiles(ctx, tx, projectId, filesToAdd); err != nil {
		return fmt.Errorf("adding some of the new project files: %w", err)
	}

	return nil
}

func (pr *projectsRepo) setFiles(ctx context.Context, tx *sql.Tx, projectId int, files []models.CodeFile, dbFiles dao.CodeFileSlice) error {
	if _, err := dbFiles.DeleteAll(ctx, tx); err != nil {
		return fmt.Errorf("deleting all existing project files: %w", err)
	}

	if err := pr.addFiles(ctx, tx, projectId, files); err != nil {
		return fmt.Errorf("adding all the new project files: %w", err)
	}

	return nil
}
