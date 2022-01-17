package models

type ProjectsList struct {
	CommonList
	Data []ProjectItem `json:"data"`
}

type ProjectItem struct {
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	UpdatedAt   int               `json:"updatedAt"`
	Files       []ProjectItemFile `json:"files"`
	Tags        []TagType         `json:"tags"`
}

type ProjectItemFile struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
