package bookhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h *BookHandler) ListBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var err error

	var page = 0

	pageQuery := r.URL.Query().Get("page")

	if pageQuery != "" {
		page, err = strconv.Atoi(pageQuery)
		if err != nil {
			h.logger.Error(fmt.Sprintf("error converting page to int: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if page < 0 {
		h.logger.Error(fmt.Sprintf("page is less than 0: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	books, err := h.store.List(page)
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

	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(booksJSON); err != nil {
		h.logger.Error(fmt.Sprintf("error writing books JSON: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}
}
