package store

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"time"

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

// New ...
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

func (pr *projectsRepo) Get(ctx context.Context, id int) (models.Project, error) {
	var p models.Project

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := pr.db.Query(queries.QueryGet, id)
	if err != nil {
		if err := ctx.Err(); err != nil {
			pr.l.Error("executing query to get project", err)
			return p, projects.ErrProjectTimeout
		}
		pr.l.Error("executing query to get project", err)
		return p, err
	}
	defer rows.Close()

	psDao := make(dao.ProjectsDao)
	for rows.Next() {
		var cf models.CodeFile
		var tag models.TagType
		var pId, cfId string

		err := rows.Scan(&pId, &p.Name, &p.Description, &cfId, &cf.Name, &cf.Content, &tag)
		if err != nil {
			return p, err
		}
		psDao.Add(pId, p, cfId, cf, tag)
	}

	if err := rows.Err(); err != nil {
		return p, err
	}

	ps := psDao.Get()
	if len(ps) == 1 {
		return ps[0], nil
	}
	return p, projects.ErrProjectNotFound
}

func (pr *projectsRepo) GetAll(ctx context.Context, q string, offset, limit int) (models.ProjectsList, error) {
	var psList models.ProjectsList

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	piOut, piErr := pr.asyncGetProjectItems(ctx, q, offset, limit)
	cOut, cErr := pr.asyncGetCount(ctx)

	if err, open := <-piErr; open {
		pr.l.Error("getting project items", err)
		return psList, err
	}

	if err, open := <-cErr; open {
		pr.l.Error("counting project records", err)
		return psList, err
	}

	psList.Data = <-piOut
	psList.TotalItems = <-cOut
	psList.Page = offset + 1
	psList.Count = len(psList.Data)
	psList.TotalPages = int(math.Ceil(float64(psList.Count) / float64(limit)))

	return psList, nil
}

func (pr *projectsRepo) asyncGetProjectItems(ctx context.Context, q string, offset, limit int) (chan []models.ProjectItem, chan error) {
	out := make(chan []models.ProjectItem, 1)
	errs := make(chan error, 1)

	go func() {
		defer func() {
			close(out)
			close(errs)
		}()

		rows, err := pr.db.Query(queries.QueryGetAll, q, limit, offset)
		if err != nil {
			if err := ctx.Err(); err != nil {
				errs <- err
				return
			}
			errs <- err
			return
		}
		defer rows.Close()

		pisDao := make(dao.ProjectItemsDao)
		for rows.Next() {
			var tag models.TagType
			var pItem models.ProjectItem
			var cfId, cfName string

			err := rows.Scan(&pItem.Id, &pItem.Name, &pItem.Description, &pItem.UpdatedAt, &cfId, &cfName, &tag)
			if err != nil {
				errs <- err
				return
			}

			pisDao.Add(pItem.Id, pItem, cfId, cfName, tag)
		}

		if err := rows.Err(); err != nil {
			errs <- err
			return
		}
		out <- pisDao.Get()
	}()
	return out, errs
}

func (pr *projectsRepo) asyncGetCount(ctx context.Context) (chan int, chan error) {
	out := make(chan int, 1)
	errs := make(chan error, 1)

	go func() {
		defer func() {
			close(out)
			close(errs)
		}()

		row := pr.db.QueryRowContext(ctx, queries.QueryCount)
		if err := row.Err(); err != nil {
			errs <- err
			return
		}
		var count int
		row.Scan(&count)
		out <- count
	}()

	return out, errs
}
