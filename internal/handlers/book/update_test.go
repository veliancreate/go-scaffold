package bookhandler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/veliancreate/books-api/internal/middleware"
	"github.com/veliancreate/books-api/internal/store"
)

const updateBody = `
	{
		"title": "Lord of the Rings"
	}
`

func TestUpdateBooksHandler(t *testing.T) {

	store := store.NewInMemoryBookStore()

	logger := middleware.NewLogger()

	handler := NewBookHandler(store, logger)

	router := httprouter.New()

	router.PATCH("/books/:id", handler.Update)

	req, err := http.NewRequest("PATCH", "/books/9dc991cf-4f73-4472-8275-e82089ca9edd", bytes.NewBuffer([]byte(updateBody)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
