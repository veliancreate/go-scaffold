package entity

import (
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Authors     []Author  `json:"authors" validate:"required"`
	Publisher   Publisher `json:"publisher" validate:"required"`
	PublishedAt string    `json:"published_at" validate:"required"`
	Pages       int       `json:"pages" validate:"required"`
}

func (p *Book) Validate() error {
	validate := validator.New()

	return validate.Struct(p)
}
