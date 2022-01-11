package projects

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

type TagType string

const (
	TagTypeUnknown      TagType = "UNKNOWN"
	TagTypeLanguage     TagType = "LANGUAGE"
	TagTypeArchitecture TagType = "ARCHITECTURE"
)
