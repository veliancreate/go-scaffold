package file_store

import (
	"context"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/veliancreate/books-api/internal/entity"
	"github.com/veliancreate/books-api/internal/store"
)

type FileStore struct {
	ioManager FileIOManager
}

func NewFileStore() *FileStore {
	ioManager := NewFileIO("seeding/books.json")
	return &FileStore{
		ioManager,
	}
}

type ByPublishedAt []entity.Book

func (b ByPublishedAt) Len() int { return len(b) }

func (b ByPublishedAt) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func (b ByPublishedAt) Less(i, j int) bool {
	patOne, _ := time.Parse(time.RFC822, b[i].PublishedAt)

	patTwo, _ := time.Parse(time.RFC822, b[j].PublishedAt)

	return patOne.Before(patTwo)
}

const limit = 10

func (fs *FileStore) List(ctx context.Context, page int) (entity.ListResponse, error) {
	var response = entity.ListResponse{}

	books, err := fs.ioManager.ReadFile()
	if err != nil {
		return response, err
	}

	response.TotalCount = len(books)

	if len(books) < limit || page == 0 {
		response.Books = books
		return response, nil
	}

	sort.Sort(ByPublishedAt(books))

	first := limit*page - limit

	last := first + limit

	if first >= len(books) {
		response.Books = []entity.Book{}
		return response, nil
	}

	if last >= len(books) {
		response.Books = books[first : len(books)-1]
		return response, nil
	}

	response.Books = books[first:last]
	return response, nil
}

func (fs *FileStore) Create(ctx context.Context, bookToCreate entity.Book) (*entity.Book, error) {
	books, err := fs.ioManager.ReadFile()
	if err != nil {
		return &bookToCreate, err
	}

	bookToCreate.ID = uuid.New()

	books = append(books, bookToCreate)

	err = fs.ioManager.WriteFile(books)
	if err != nil {
		return &bookToCreate, err
	}

	return &bookToCreate, nil
}

func (fs *FileStore) Update(ctx context.Context, id uuid.UUID, bookUpdateDetails entity.Book) (*entity.Book, error) {
	var book *entity.Book

	books, err := fs.ioManager.ReadFile()
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

	err = fs.ioManager.WriteFile(books)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (fs *FileStore) Delete(ctx context.Context, id uuid.UUID) error {
	books, err := fs.ioManager.ReadFile()
	if err != nil {
		return err
	}

	newBooks := []entity.Book{}

	for i := 0; i < len(books); i++ {
		if books[i].ID != id {
			newBooks = append(newBooks, books[i])
		}
	}

	err = fs.ioManager.WriteFile(newBooks)
	if err != nil {
		return err
	}

	return nil
}
