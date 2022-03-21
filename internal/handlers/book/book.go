package bookhandler

import (
	"github.com/google/uuid"
	"github.com/veliancreate/books-api/internal/entity"
)

type BookStore interface {
	List(page int) (entity.ListResponse, error)
	Create(bookDetails entity.Book) (*entity.Book, error)
	Delete(id uuid.UUID) error
	Update(id uuid.UUID, bookDetails entity.Book) (*entity.Book, error)
}

type Logger interface {
	Error(message string)
	Info(message string)
}

type BookHandler struct {
	store  BookStore
	logger Logger
}

func NewBookHandler(store BookStore, logger Logger) *BookHandler {
	return &BookHandler{
		store,
		logger,
	}
}
