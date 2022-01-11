package validate

import "github.com/go-playground/validator/v10"

var v = validator.New()

func Get() *validator.Validate {
	return v
}
