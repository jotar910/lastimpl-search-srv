package dao

import (
	"lastimplementation.com/pkg/services/projects/models"
)

type ProjectItemsDao map[string]ProjectItemDao

func (pisDao ProjectItemsDao) Add(id string, project models.ProjectItem, cfId string, cfName string, tag models.TagType) {
	p, ok := pisDao.Has(id)
	if !ok {
		p = newProjectItemDao(id, project)
		pisDao[id] = p
	}
	p.tags.Add(tag)
	p.files.Add(cfId, cfName)
}

func (pisDao ProjectItemsDao) Get() []models.ProjectItem {
	ps := make([]models.ProjectItem, 0)
	for _, piDao := range pisDao {
		ps = append(ps, piDao.Get())
	}
	return ps
}

func (pisDao ProjectItemsDao) Has(id string) (ProjectItemDao, bool) {
	ps, ok := pisDao[id]
	return ps, ok
}

type ProjectItemDao struct {
	id      string
	tags    TagsDao
	files   ProjectItemFilesDao
	project models.ProjectItem
}

func newProjectItemDao(id string, project models.ProjectItem) ProjectItemDao {
	return ProjectItemDao{
		id:      id,
		tags:    make(TagsDao),
		files:   make(ProjectItemFilesDao),
		project: project,
	}
}

func (piDao ProjectItemDao) Get() models.ProjectItem {
	piDao.project.Files = piDao.files.Get()
	piDao.project.Tags = piDao.tags.Get()
	return piDao.project
}

type ProjectItemFilesDao map[string]ProjectItemFileDao

func (pifsDao ProjectItemFilesDao) Add(id string, name string) {
	if _, ok := pifsDao.Has(id); !ok {
		pifsDao[id] = newProjectItemFileDao(id, name)
	}
}

func (pifsDao ProjectItemFilesDao) Get() []models.ProjectItemFile {
	var ps []models.ProjectItemFile
	for _, pifDao := range pifsDao {
		ps = append(ps, pifDao.Get())
	}
	return ps
}

func (pifsDao ProjectItemFilesDao) Has(id string) (ProjectItemFileDao, bool) {
	ps, ok := pifsDao[id]
	return ps, ok
}

type ProjectItemFileDao struct {
	id   string
	name string
}

func newProjectItemFileDao(id string, name string) ProjectItemFileDao {
	return ProjectItemFileDao{id, name}
}

func (pifDao ProjectItemFileDao) Get() models.ProjectItemFile {
	return models.ProjectItemFile{
		Id:   pifDao.id,
		Name: pifDao.name,
	}
}
