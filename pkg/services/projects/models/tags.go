package models

type TagType string

const (
	TagTypeUnknown      TagType = "UNKNOWN"
	TagTypeLanguage     TagType = "LANGUAGE"
	TagTypeArchitecture TagType = "ARCHITECTURE"
)
