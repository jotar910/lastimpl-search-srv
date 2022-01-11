package dao

import "lastimplementation.com/pkg/services/projects"

type ProjectsDao map[int]ProjectDao

func (psDao ProjectsDao) Add(id int, project projects.Project, cfId int, cf projects.CodeFile, tag projects.TagType) {
	p, ok := psDao.Has(id)
	if !ok {
		p = newProjectDao(id, project)
		psDao[id] = p
	}
	p.tags.Add(tag)
	p.codeFiles.Add(cfId, cf)
}

func (psDao ProjectsDao) Get() []projects.Project {
	var ps []projects.Project
	for _, pDao := range psDao {
		ps = append(ps, pDao.Get())
	}
	return ps
}

func (psDao ProjectsDao) Has(id int) (ProjectDao, bool) {
	ps, ok := psDao[id]
	return ps, ok
}

type ProjectDao struct {
	id        int
	tags      TagsDao
	codeFiles CodeFilesDao
	project   projects.Project
}

func newProjectDao(id int, project projects.Project) ProjectDao {
	return ProjectDao{
		id:        id,
		tags:      make(TagsDao),
		codeFiles: make(CodeFilesDao),
		project:   project,
	}
}

func (pDao ProjectDao) Get() projects.Project {
	pDao.project.Tag = pDao.tags.Get()
	pDao.project.Files = pDao.codeFiles.Get()
	return pDao.project
}

type CodeFilesDao map[int]CodeFileDao

func (cfsDao CodeFilesDao) Add(id int, codeFile projects.CodeFile) {
	if _, ok := cfsDao.Has(id); !ok {
		cfsDao[id] = newCodeFileDao(id, codeFile)
	}
}

func (cfsDao CodeFilesDao) Has(id int) (CodeFileDao, bool) {
	cf, ok := cfsDao[id]
	return cf, ok
}

func (cfsDao CodeFilesDao) Get() []projects.CodeFile {
	var cfs []projects.CodeFile
	for _, cfDao := range cfsDao {
		cfs = append(cfs, cfDao.Get())
	}
	return cfs
}

type CodeFileDao struct {
	id       int
	codeFile projects.CodeFile
}

func newCodeFileDao(id int, codeFile projects.CodeFile) CodeFileDao {
	return CodeFileDao{
		id:       id,
		codeFile: codeFile,
	}
}

func (cfDao CodeFileDao) Get() projects.CodeFile {
	return cfDao.codeFile
}

type TagsDao map[projects.TagType]struct{}

func (tagsDao TagsDao) Add(tag projects.TagType) {
	tagsDao[tag] = struct{}{}
}

func (tagsDao TagsDao) Get() []projects.TagType {
	var ts []projects.TagType
	for t := range tagsDao {
		ts = append(ts, t)
	}
	return ts
}
