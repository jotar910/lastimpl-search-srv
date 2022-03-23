package models

import (
	"encoding/json"
	"io"
)

const (
	MaximumCodeFiles int = 50
)

type ProjectDetails struct {
	Name        string `json:"name" validate:"min=5,max=50"`
	Description string `json:"description" validate:"min=20,max=500"`
}

func (pd *ProjectDetails) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(pd)
}

type Project struct {
	ProjectDetails
	Id    int        `json:"id,omitempty"`
	Tag   []TagType  `json:"tags" validate:"max=20"`
	Files []CodeFile `json:"files" validate:"max=50"`
}

func (p *Project) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(p)
}

func (p *Project) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(p)
}

type CodeFile struct {
	Id      int    `json:"id,omitempty"`
	Name    string `json:"name" validate:"max=50"`
	Content string `json:"content" validate:"max=10000"`
}

type CodeFiles []CodeFile

func (cfs *CodeFiles) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(cfs)
}

func (cfs *CodeFiles) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(cfs)
}
