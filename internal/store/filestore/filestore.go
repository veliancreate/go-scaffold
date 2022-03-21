package store

import (
	"github.com/google/uuid"
	"github.com/veliancreate/books-api/internal/entity"
	"github.com/veliancreate/books-api/internal/store"
)

type FileStore struct {
}

func NewFileStore() *FileStore {
	return &FileStore{}
}

func (fs *FileStore) List() ([]entity.Book, error) {
	return getBooks()
}

func (fs *FileStore) Create(bookToCreate entity.Book) (*entity.Book, error) {
	books, err := getBooks()
	if err != nil {
		return &bookToCreate, err
	}

	bookToCreate.ID = uuid.New()

	books = append(books, bookToCreate)

	err = writeFile(books)
	if err != nil {
		return &bookToCreate, err
	}

	return &bookToCreate, nil
}

func (fs *FileStore) Update(id uuid.UUID, bookUpdateDetails entity.Book) (*entity.Book, error) {
	var book *entity.Book

	books, err := getBooks()
	if err != nil {
		return book, err
	}

	for i := 0; i < len(books); i++ {
		if books[i].ID == id {
			book = &books[i]
		}
	}

	if book == nil {
		return book, store.ErrNotFound
	}

	if bookUpdateDetails.Title != "" {
		book.Title = bookUpdateDetails.Title
	}

	if bookUpdateDetails.PublishedAt != "" {
		book.PublishedAt = bookUpdateDetails.PublishedAt
	}

	if bookUpdateDetails.Pages != 0 {
		book.Pages = bookUpdateDetails.Pages
	}

	err = writeFile(books)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (fs *FileStore) Delete(id uuid.UUID) error {
	books, err := getBooks()
	if err != nil {
		return err
	}

	newBooks := []entity.Book{}

	for i := 0; i < len(books); i++ {
		if books[i].ID != id {
			newBooks = append(newBooks, books[i])
		}
	}

	err = writeFile(newBooks)
	if err != nil {
		return err
	}

	return nil
}
