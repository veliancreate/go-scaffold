package bookhandler

import (
	"context"

	"github.com/google/uuid"
	"github.com/veliancreate/books-api/internal/entity"
)

type BookStore interface {
	List(ctx context.Context) ([]entity.Book, error)
	Create(ctx context.Context, book entity.Book) (entity.Book, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, bookDetails entity.Book) (entity.Book, error)
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
