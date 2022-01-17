package dao

import (
	"lastimplementation.com/pkg/services/projects/models"
)

type ProjectsDao map[string]ProjectDao

func (psDao ProjectsDao) Add(id string, project models.Project, cfId string, cf models.CodeFile, tag models.TagType) {
	p, ok := psDao.Has(id)
	if !ok {
		p = newProjectDao(id, project)
		psDao[id] = p
	}
	p.tags.Add(tag)
	p.codeFiles.Add(cfId, cf)
}

func (psDao ProjectsDao) Get() []models.Project {
	var ps []models.Project
	for _, pDao := range psDao {
		ps = append(ps, pDao.Get())
	}
	return ps
}

func (psDao ProjectsDao) Has(id string) (ProjectDao, bool) {
	ps, ok := psDao[id]
	return ps, ok
}

type ProjectDao struct {
	id        string
	tags      TagsDao
	codeFiles CodeFilesDao
	project   models.Project
}

func newProjectDao(id string, project models.Project) ProjectDao {
	return ProjectDao{
		id:        id,
		tags:      make(TagsDao),
		codeFiles: make(CodeFilesDao),
		project:   project,
	}
}

func (pDao ProjectDao) Get() models.Project {
	pDao.project.Tag = pDao.tags.Get()
	pDao.project.Files = pDao.codeFiles.Get()
	return pDao.project
}

type CodeFilesDao map[string]CodeFileDao

func (cfsDao CodeFilesDao) Add(id string, codeFile models.CodeFile) {
	if _, ok := cfsDao.Has(id); !ok {
		cfsDao[id] = newCodeFileDao(id, codeFile)
	}
}

func (cfsDao CodeFilesDao) Has(id string) (CodeFileDao, bool) {
	cf, ok := cfsDao[id]
	return cf, ok
}

func (cfsDao CodeFilesDao) Get() []models.CodeFile {
	var cfs []models.CodeFile
	for _, cfDao := range cfsDao {
		cfs = append(cfs, cfDao.Get())
	}
	return cfs
}

type CodeFileDao struct {
	id       string
	codeFile models.CodeFile
}

func newCodeFileDao(id string, codeFile models.CodeFile) CodeFileDao {
	return CodeFileDao{
		id:       id,
		codeFile: codeFile,
	}
}

func (cfDao CodeFileDao) Get() models.CodeFile {
	return cfDao.codeFile
}
