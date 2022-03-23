package projects

import (
	"context"

	"lastimplementation.com/pkg/services/projects/logger"
	"lastimplementation.com/pkg/services/projects/models"
)

type Repo interface {
	Reset(ctx context.Context) error
	Get(ctx context.Context, id int) (models.Project, error)
	GetAll(ctx context.Context, q string, page, limit int) (models.ProjectsList, error)
	Add(ctx context.Context, project models.Project) (int, error)
	Update(ctx context.Context, id int, details models.ProjectDetails) error
	Delete(ctx context.Context, id int) error
	UpdateFiles(ctx context.Context, projectId int, files []models.CodeFile) error
	GetFiles(ctx context.Context, projectId int) (models.CodeFiles, error)
}

type Service interface {
	ResetRepo(ctx context.Context) error
	Get(ctx context.Context, id int) (models.Project, error)
	GetAll(ctx context.Context, qp models.SearchQP) (models.ProjectsList, error)
	Add(ctx context.Context, project models.Project) (int, error)
	Update(ctx context.Context, id int, details models.ProjectDetails) error
	Delete(ctx context.Context, id int) error
	GetFiles(ctx context.Context, projectId int) (models.CodeFiles, error)
	UpdateFiles(ctx context.Context, projectId int, files []models.CodeFile) error
}

type projects struct {
	l    logger.Logger
	repo Repo
}

// New creates a new projects service.
func New(l logger.Logger, repo Repo) *projects {
	return &projects{
		l:    l.WithPrefix("service"),
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
	return p.repo.GetAll(ctx, qp.Query, qp.Page, qp.Limit)
}

// Add adds a new project.
func (p *projects) Add(ctx context.Context, project models.Project) (int, error) {
	return p.repo.Add(ctx, project)
}

// Update updates an existing project.
func (p *projects) Update(ctx context.Context, id int, details models.ProjectDetails) error {
	return p.repo.Update(ctx, id, details)
}

// Delete deletes an existing project.
func (p *projects) Delete(ctx context.Context, id int) error {
	return p.repo.Delete(ctx, id)
}

// GetFiles returns a list of the files that exists on a project.
func (p *projects) GetFiles(ctx context.Context, projectId int) (models.CodeFiles, error) {
	return p.repo.GetFiles(ctx, projectId)
}

// UpdateFiles updates the code files on a project.
func (p *projects) UpdateFiles(ctx context.Context, projectId int, files []models.CodeFile) error {
	return p.repo.UpdateFiles(ctx, projectId, files)
}
