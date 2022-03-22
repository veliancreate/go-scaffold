package in_memory_store

import (
	"context"

	"github.com/google/uuid"
	"github.com/veliancreate/books-api/internal/entity"
	"github.com/veliancreate/books-api/internal/store"
)

type InMemoryBookStore struct {
	books []entity.Book
}

func NewInMemoryBookStore() *InMemoryBookStore {
	id, _ := uuid.Parse("9dc991cf-4f73-4472-8275-e82089ca9edd")
	return &InMemoryBookStore{
		books: []entity.Book{
			{
				ID:    id,
				Title: "Lord of the Rings",
				Authors: []entity.Author{
					{
						Name: "J.R.R Tolkein",
					},
				},
				Publisher: entity.Publisher{
					Name: "Penguin",
				},
				PublishedAt: "1st January 1900",
				Pages:       1000,
			},
		},
	}
}

func (bs *InMemoryBookStore) List(ctx context.Context, page int) (entity.ListResponse, error) {
	var response = entity.ListResponse{
		TotalCount: len(bs.books),
		Books:      bs.books,
	}

	return response, nil
}

func (bs *InMemoryBookStore) Update(ctx context.Context, id uuid.UUID, bookUpdateDetails entity.Book) (*entity.Book, error) {
	var book *entity.Book

	for i := 0; i < len(bs.books); i++ {
		if bs.books[i].ID == id {
			book = &bs.books[i]
		}
	}

	if book == nil {
		return book, store.ErrNotFound
	}

	book.Authors = bookUpdateDetails.Authors

	book.Title = bookUpdateDetails.Title

	book.PublishedAt = bookUpdateDetails.Title

	book.Publisher = bookUpdateDetails.Publisher

	book.Pages = bookUpdateDetails.Pages

	return book, nil
}

func (bs *InMemoryBookStore) Create(ctx context.Context, bookUpdateDetails entity.Book) (*entity.Book, error) {
	bs.books = append(bs.books, bookUpdateDetails)

	return &bookUpdateDetails, nil
}

func (bs *InMemoryBookStore) Delete(ctx context.Context, id uuid.UUID) error {
	var newBooks []entity.Book

	for i := 0; i < len(bs.books); i++ {
		if bs.books[i].ID != id {
			newBooks = append(newBooks, bs.books[i])
		}
	}

	bs.books = newBooks

	return nil
}
