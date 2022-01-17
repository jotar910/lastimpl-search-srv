package projects

import (
	"context"

	"lastimplementation.com/pkg/services/projects/logger"
	"lastimplementation.com/pkg/services/projects/models"
)

type Repo interface {
	Reset(ctx context.Context) error
	Get(ctx context.Context, id int) (models.Project, error)
	GetAll(ctx context.Context, q string, offset, limit int) (models.ProjectsList, error)
}

type Service interface {
	ResetRepo(ctx context.Context) error
	GetAll(ctx context.Context, qp models.SearchQP) (models.ProjectsList, error)
	Get(ctx context.Context, id int) (models.Project, error)
}

type projects struct {
	l    logger.Logger
	repo Repo
}

// New ...
func New(l logger.Logger, repo Repo) Service {
	return &projects{
		l:    l,
		repo: repo,
	}
}

// Reset resets the projects repo.
func (p *projects) ResetRepo(ctx context.Context) error {
	return p.repo.Reset(ctx)
}

// Get gets a single project.
func (p *projects) Get(ctx context.Context, id int) (models.Project, error) {
	return p.repo.Get(ctx, id)
}

// GetAll gets all the projects.
func (p *projects) GetAll(ctx context.Context, qp models.SearchQP) (models.ProjectsList, error) {
	return p.repo.GetAll(ctx, qp.Query, qp.Page-1, qp.Limit)
}
