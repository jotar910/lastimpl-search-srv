package projects

import "errors"

var (
	ErrProjectNotFound = errors.New("requested project could not be found")
)
