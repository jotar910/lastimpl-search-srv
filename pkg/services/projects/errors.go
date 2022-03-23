package projects

import (
	"encoding/json"
	"errors"
	"io"
)

var (
	ErrProjectTimeout           = NewError("request timeout")
	ErrProjectNotFound          = NewError("requested project could not be found")
	ErrAddProjectDuplicatedName = NewError("duplicated name")
	ErrDecodeBody               = NewError("failed to decode body")
)

type outboundError struct {
	Message string `json:"message"`
}

type OutboundError struct {
	e error
}

func NewError(message string) OutboundError {
	return OutboundError{errors.New(message)}
}

func (err OutboundError) Error() string {
	return err.e.Error()
}

func (err OutboundError) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(outboundError{err.Error()})
}
