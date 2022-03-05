package models

type TagType string

const (
	TagTypeUnknown      TagType = "UNKNOWN"
	TagTypeLanguage     TagType = "LANGUAGE"
	TagTypeArchitecture TagType = "ARCHITECTURE"
)

type Tag struct {
	Id   int     `json:"id"`
	Name TagType `json:"name"`
}
