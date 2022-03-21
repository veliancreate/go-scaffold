package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/veliancreate/books-api/internal/entity"
)

type InMemoryBookStore struct {
	books []entity.Book
}

func NewInMemoryBookStore() *InMemoryBookStore {
	return &InMemoryBookStore{
		books: []entity.Book{
			{
				ID:    uuid.New(),
				Title: "Lord of the Rings",
				Authors: []entity.Author{
					{
						ID:   uuid.New(),
						Name: "J.R.R Tolkein",
					},
				},
				Publisher: entity.Publisher{
					ID:   uuid.New(),
					Name: "Penguin",
				},
				PublishedAt: "1st January 1900",
				Pages:       1000,
			},
		},
	}
}

func (bs *InMemoryBookStore) List(ctx context.Context) ([]entity.Book, error) {
	return bs.books, nil
}

func (bs *InMemoryBookStore) Update(ctx context.Context, id uuid.UUID, bookUpdateDetails entity.Book) (entity.Book, error) {
	var book *entity.Book

	for i := 0; i < len(bs.books); i++ {
		if bs.books[i].ID == id {
			book = &bs.books[i]
		}
	}

	if book == nil {
		return *book, ErrNotFound
	}

	book.Authors = bookUpdateDetails.Authors

	book.Title = bookUpdateDetails.Title

	book.PublishedAt = bookUpdateDetails.Title

	book.Publisher = bookUpdateDetails.Publisher

	book.Pages = bookUpdateDetails.Pages

	return *book, nil
}

func (bs *InMemoryBookStore) Create(ctx context.Context, bookUpdateDetails entity.Book) (entity.Book, error) {
	bs.books = append(bs.books, bookUpdateDetails)

	return bookUpdateDetails, nil
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
