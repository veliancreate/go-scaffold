package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/veliancreate/books-api/internal/entity"
)

type BookStore interface {
	List() ([]entity.Book, error)
	Create(bookDetails entity.Book) (entity.Book, error)
	Delete(id uuid.UUID) error
	Update(id uuid.UUID, bookDetails entity.Book) (entity.Book, error)
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

func (h *BookHandler) ListBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	books, err := h.store.List()
	if err != nil {
		h.logger.Error(fmt.Sprintf("error listing books: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	booksJSON, err := json.Marshal(books)
	if err != nil {
		h.logger.Error(fmt.Sprintf("error marshalling books JSON: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("successfully returned books")
	w.Write(booksJSON)
}

func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.Error(fmt.Sprintf("could not read body: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var bookToCreate *entity.Book

	err = json.Unmarshal(body, bookToCreate)
	if err != nil {
		h.logger.Error(fmt.Sprintf("could not marshal create book input: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}

	createdBook, err := h.store.Create(*bookToCreate)
	if err != nil {
		h.logger.Error(fmt.Sprintf("could not create book: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bookJSON, err := json.Marshal(createdBook)
	if err != nil {
		h.logger.Error(fmt.Sprintf("error marshalling books JSON: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("successfully created book")
	w.Write(bookJSON)
}

func (h *BookHandler) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		h.logger.Error(fmt.Sprintf("error parsing uuid %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.store.Delete(parsedID)
	if err != nil {
		h.logger.Error(fmt.Sprintf("error deleting book %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *BookHandler) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.Error(fmt.Sprintf("could not read body: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var updateDetails *entity.Book

	err = json.Unmarshal(body, updateDetails)
	if err != nil {
		h.logger.Error(fmt.Sprintf("could not marshal create book input: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}

	updatedBook, err := h.store.Create(*updateDetails)
	if err != nil {
		h.logger.Error(fmt.Sprintf("could not update book: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bookJSON, err := json.Marshal(updatedBook)
	if err != nil {
		h.logger.Error(fmt.Sprintf("error marshalling books JSON: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("successfully updated book")
	w.Write(bookJSON)
}
