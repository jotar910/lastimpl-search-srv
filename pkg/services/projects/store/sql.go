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
	return &projectsRepo{l, db}
}

// Reset resets the projects tables.
func (pr *projectsRepo) Reset(ctx context.Context) error {
	pr.l.Trace("Resetting the projects repo")
	if _, err := pr.db.Exec(queries.QueryDropTables); err != nil {
		return fmt.Errorf("executing query: %w", err)
	}
	if _, err := pr.db.Exec(queries.QueryCreateTables); err != nil {
		return fmt.Errorf("executing query: %w", err)
	}
	return nil
}

func (pr *projectsRepo) Add(ctx context.Context, project models.Project) (int, error) {
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		pr.l.Error("begining transaction", err.Error())
		return -1, err
	}

	p := dao.Project{Name: project.Name, Description: project.Description}
	if err = p.Insert(ctx, tx, boil.Infer()); err != nil {
		pr.l.Error("inserting new project", err)
		tx.Rollback()
		if strings.HasSuffix(err.Error(), ErrDuplicated("projects_name_key")) {
			return -1, projects.ErrAddProjectDuplicatedName
		}
		return -1, err
	}

	if err := pr.addTags(ctx, tx, p.ID, project.Tag); err != nil {
		pr.l.Error("adding tags to the project", err.Error())
		tx.Rollback()
		return -1, err
	}

	if err := pr.addFiles(ctx, tx, p.ID, project.Files); err != nil {
		pr.l.Error("adding code files to the project", err.Error())
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
			return err
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

func (pr *projectsRepo) Get(ctx context.Context, id int) (models.Project, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		pr.l.Error("begining transaction", err)
		return models.Project{}, err
	}

	p, err := dao.Projects(
		qm.Where("id = ?", id),
		qm.Select(dao.ProjectColumns.ID, dao.ProjectColumns.Name, dao.ProjectColumns.Description),
		qm.Load(dao.ProjectRels.CodeFiles,
			qm.Select(dao.CodeFileColumns.ID, dao.CodeFileColumns.Name, dao.CodeFileColumns.Content)),
		qm.Load(qm.Rels(dao.ProjectRels.ProjectsTags), qm.Select(dao.ProjectsTagColumns.TagID)),
		qm.Load(qm.Rels(dao.ProjectRels.ProjectsTags, dao.ProjectsTagRels.Tag), qm.Select(dao.TagColumns.ID, dao.TagColumns.Name)),
	).One(ctx, tx)
	if err != nil {
		pr.l.Error("executing query to get project:", err)
		if strings.HasSuffix(err.Error(), ErrNotResult()) {
			return models.Project{}, projects.ErrProjectNotFound
		}
		return models.Project{}, err
	}

	res := models.Project{
		Id: p.ID,
		ProjectDetails: models.ProjectDetails{
			Name:        p.Name,
			Description: p.Description,
		},
	}

	for _, tag := range p.R.ProjectsTags {
		res.Tag = append(res.Tag, models.TagType(tag.R.Tag.Name))
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

func (pr *projectsRepo) GetAll(ctx context.Context, query string, page, limit int) (models.ProjectsList, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		pr.l.Error("begining transaction", err)
		return models.ProjectsList{}, err
	}

	projects, err := dao.Projects(
		qm.Select(dao.ProjectColumns.ID, dao.ProjectColumns.Name, dao.ProjectColumns.Description, dao.ProjectColumns.UpdatedAt),
		qm.Where("LOWER(name) ~ ?", strings.ToLower(query)),
		qm.Offset((page-1)*limit),
		qm.Limit(limit),
		qm.Load(dao.ProjectRels.CodeFiles,
			qm.Select(dao.CodeFileColumns.ProjectID, dao.CodeFileColumns.ID, dao.CodeFileColumns.Name)),
		qm.Load(dao.ProjectRels.ProjectsTags,
			qm.Select(dao.ProjectsTagColumns.ProjectID, dao.ProjectsTagColumns.TagID)),
		qm.Load(qm.Rels(dao.ProjectRels.ProjectsTags, dao.ProjectsTagRels.Tag),
			qm.Select(dao.TagColumns.ID, dao.TagColumns.Name)),
	).All(ctx, tx)
	if err != nil {
		pr.l.Error("getting project items", err)
		tx.Rollback()
		return models.ProjectsList{}, err
	}

	count, err := dao.Projects().Count(ctx, pr.db)
	if err != nil {
		pr.l.Error("counting total project items", err)
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

func (pr *projectsRepo) Update(ctx context.Context, id int, project models.ProjectDetails) error {
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		pr.l.Error("begining transaction", err)
		return err
	}

	p := dao.Project{
		ID:          id,
		Name:        project.Name,
		Description: project.Description,
	}
	changed, err := p.Update(ctx, tx, boil.Whitelist(dao.ProjectColumns.Name, dao.ProjectColumns.Description, dao.ProjectColumns.UpdatedAt))
	if err != nil {
		pr.l.Error("updating existing project", err)
		tx.Rollback()
		return err
	}
	if changed == 0 {
		err := projects.ErrProjectNotFound
		pr.l.Error("updating existing project", err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (pr *projectsRepo) Delete(ctx context.Context, id int) error {
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		pr.l.Error("begining transaction", err)
		return err
	}

	p, err := dao.Projects(
		qm.Select(dao.ProjectColumns.ID),
		qm.Where("id = ?", id),
	).One(ctx, tx)
	if err != nil {
		pr.l.Error("finding project", err)
		tx.Rollback()
		return projects.ErrProjectNotFound
	}

	if _, err := dao.CodeFiles(qm.Where("project_id = ?", id)).DeleteAll(ctx, tx); err != nil {
		pr.l.Error("deleting code files from existing project", err)
		tx.Rollback()
		return err
	}

	if _, err := dao.ProjectsTags(qm.Where("project_id = ?", id)).DeleteAll(ctx, tx); err != nil {
		pr.l.Error("deleting code files from existing project", err)
		tx.Rollback()
		return err
	}

	if _, err := p.Delete(ctx, tx); err != nil {
		pr.l.Error("deleting existing project", err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
