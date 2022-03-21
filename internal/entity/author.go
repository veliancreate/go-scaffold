package entity

import (
	"github.com/go-playground/validator"
)

type Author struct {
	Name string `json:"name"`
}

func (p *Author) Validate() error {
	validate := validator.New()

	return validate.Struct(p)
}
