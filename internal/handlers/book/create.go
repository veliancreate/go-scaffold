package bookhandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/veliancreate/books-api/internal/entity"
)

func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	r.Body = http.MaxBytesReader(w, r.Body, 20000)

	dec := json.NewDecoder(r.Body)

	dec.DisallowUnknownFields()

	var bookToCreate *entity.Book

	err := dec.Decode(bookToCreate)
	if err != nil {
		h.logger.Error(fmt.Sprintf("could not read body: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = bookToCreate.Validate()
	if err != nil {
		h.logger.Error(fmt.Sprintf("book creation validation error: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
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
