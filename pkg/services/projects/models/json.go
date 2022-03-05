package models

import "io"

type EncodeJSON interface {
	ToJSON(io.Writer) error
}

type DecodeJSON interface {
	FromJSON(io.Reader) error
}
