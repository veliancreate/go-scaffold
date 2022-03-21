package entity

import (
	"github.com/go-playground/validator"
)

type Publisher struct {
	Name string `json:"name"`
}

func (p *Publisher) Validate() error {
	validate := validator.New()

	return validate.Struct(p)
}
