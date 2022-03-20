package entity

import (
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type Publisher struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name" validate:"required"`
}

func (p *Publisher) Validate() error {
	validate := validator.New()

	return validate.Struct(p)
}
