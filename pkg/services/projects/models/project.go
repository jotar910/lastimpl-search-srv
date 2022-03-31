package models

import (
	"encoding/json"
	"io"
	"reflect"
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

func (pd *ProjectDetails) FilledProps() []string {
	props := make([]string, 0)
	v := reflect.ValueOf(pd).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if field := v.Field(i); !field.IsZero() {
			props = append(props, t.Field(i).Name)
		}
	}
	return props
}

type Project struct {
	ProjectDetails
	Id    int        `json:"id,omitempty"`
	Tags  []TagType  `json:"tags" validate:"max=30"`
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
