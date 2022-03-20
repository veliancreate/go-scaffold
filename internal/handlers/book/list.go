package bookhandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

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
