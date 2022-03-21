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

var body = `
{
	"title": "Harry Potter2",
	"authors": [
			{
					"name": "J.K Rowling"
			}
	],
	"publisher": {
			"name": "Hodder"
	},
	"published_at": "1st January 1999",
	"pages": 2000
}
`

func TestCreateBooksHandler(t *testing.T) {

	store := store.NewInMemoryBookStore()

	logger := middleware.NewLogger()

	handler := NewBookHandler(store, logger)

	router := httprouter.New()

	router.POST("/books", handler.Create)

	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer([]byte(body)))
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
