package entity

import "github.com/google/uuid"

type Author struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Publisher struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Book struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Authors     []Author  `json:"authors"`
	Publisher   Publisher `json:"publisher"`
	PublishedAt string    `json:"published_at"`
	Pages       int       `json:"pages"`
}
