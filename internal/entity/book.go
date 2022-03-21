package entity

import (
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Authors     []Author  `json:"authors"`
	Publisher   Publisher `json:"publisher"`
	PublishedAt string    `json:"published_at"`
	Pages       int       `json:"pages"`
}

func (p *Book) Validate() error {
	validate := validator.New()

	return validate.Struct(p)
}
