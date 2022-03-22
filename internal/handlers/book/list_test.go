package bookhandler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/veliancreate/books-api/internal/middleware"
	store "github.com/veliancreate/books-api/internal/store/in_memory"
)

func TestListBooksHandler(t *testing.T) {

	store := store.NewInMemoryBookStore()

	logger := middleware.NewLogger()

	handler := NewBookHandler(store, logger)

	router := httprouter.New()

	router.GET("/books", handler.ListBooks)

	req, err := http.NewRequest("GET", "/books", nil)
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
