package projects

import "errors"

var (
	ErrProjectTimeout  = errors.New("request timeout")
	ErrProjectNotFound = errors.New("requested project could not be found")
)
