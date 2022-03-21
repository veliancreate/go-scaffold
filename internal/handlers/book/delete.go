package bookhandler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

func (h *BookHandler) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		h.logger.Error(fmt.Sprintf("error parsing uuid %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.store.Delete(r.Context(), parsedID)
	if err != nil {
		h.logger.Error(fmt.Sprintf("error deleting book %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.logger.Info("book deleted")

	w.WriteHeader(http.StatusNoContent)
}
