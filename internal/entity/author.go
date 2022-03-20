package entity

import (
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type Author struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name" validate:"required"`
}

func (p *Author) Validate() error {
	validate := validator.New()

	return validate.Struct(p)
}
