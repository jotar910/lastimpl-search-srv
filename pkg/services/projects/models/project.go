package models

type Project struct {
	Name        string
	Description string
	Tag         []TagType
	Files       []CodeFile
}

type CodeFile struct {
	Name    string
	Content string
}
