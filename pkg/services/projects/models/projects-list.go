package models

import (
	"encoding/json"
	"io"
)

type ProjectsList struct {
	CommonList
	Data []ProjectItem `json:"data"`
}

func (pl *ProjectsList) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(pl)
}

type ProjectItem struct {
	Id          int               `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	UpdatedAt   int64             `json:"updatedAt"`
	Files       []ProjectItemFile `json:"files"`
	Tags        []Tag             `json:"tags"`
}

type ProjectItemFile struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
