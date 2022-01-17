package dao

import (
	"lastimplementation.com/pkg/services/projects/models"
)

type TagsDao map[models.TagType]struct{}

func (tagsDao TagsDao) Add(tag models.TagType) {
	tagsDao[tag] = struct{}{}
}

func (tagsDao TagsDao) Get() []models.TagType {
	var ts []models.TagType
	for t := range tagsDao {
		ts = append(ts, t)
	}
	return ts
}
