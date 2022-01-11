package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"lastimplementation.com/pkg/services/projects"
	"lastimplementation.com/pkg/services/projects/logger"
	"lastimplementation.com/pkg/services/projects/store/dao"
	"lastimplementation.com/pkg/services/projects/store/queries"
)

var (
	errorQuery = errors.New("unable to query the database")
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

func (pr *projectsRepo) Get(ctx context.Context, id int) (projects.Project, error) {
	var p projects.Project

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := pr.db.Query(queries.QueryGet, id)
	if err != nil {
		if err := ctx.Err(); err != nil {
			pr.l.Error("executing query to get project", err)
			return p, err
		}
		pr.l.Error("executing query to get project", err)
		return p, errorQuery
	}
	defer rows.Close()

	ps, err := readProjectRows(rows)
	if err != nil {
		return p, err
	}
	if len(ps) == 1 {
		return ps[0], nil
	}
	return p, projects.ErrProjectNotFound
}

func (pr *projectsRepo) GetAll(ctx context.Context, q string, offset, limit int) ([]projects.Project, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := pr.db.Query(queries.QueryGetAll, q, limit, offset)
	if err != nil {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		return nil, errorQuery
	}
	defer rows.Close()

	return readProjectRows(rows)
}

func readProjectRows(rows *sql.Rows) ([]projects.Project, error) {
	psDao := make(dao.ProjectsDao)
	for rows.Next() {
		var p projects.Project
		var cf projects.CodeFile
		var tag projects.TagType
		var pId, cfId int

		err := rows.Scan(&pId, &p.Name, &p.Description, &cfId, &cf.Name, &cf.Content, &tag)
		if err != nil {
			return nil, err
		}
		psDao.Add(pId, p, cfId, cf, tag)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return psDao.Get(), nil
}
