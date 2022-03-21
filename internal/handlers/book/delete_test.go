package bookhandler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/veliancreate/books-api/internal/middleware"
	"github.com/veliancreate/books-api/internal/store"
)

func TestDeleteBooksHandler(t *testing.T) {

	store := store.NewInMemoryBookStore()

	logger := middleware.NewLogger()

	handler := NewBookHandler(store, logger)

	router := httprouter.New()

	router.DELETE("/books/:id", handler.Delete)

	req, err := http.NewRequest("DELETE", "/books/97fa23d8-524e-489d-9a66-c9d53c846bb7", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
