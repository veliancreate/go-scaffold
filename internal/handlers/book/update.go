package bookhandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/veliancreate/books-api/internal/entity"
	"github.com/veliancreate/books-api/internal/handlers"
)

func (h *BookHandler) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, 20000)

	dec := json.NewDecoder(r.Body)

	dec.DisallowUnknownFields()

	var updateDetails *entity.Book

	err := dec.Decode(&updateDetails)
	if err != nil {
		handlers.HandleJSONParsingError(err, w, h.logger)
		return
	}

	err = updateDetails.Validate()
	if err != nil {
		h.logger.Error(fmt.Sprintf("book update validation error: %v", err))
		http.Error(w, "required field missing", http.StatusBadRequest)
		return
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

	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(bookJSON); err != nil {
		h.logger.Error(fmt.Sprintf("error writing books JSON: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}
}
