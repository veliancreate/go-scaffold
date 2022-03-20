package server

import (
	"github.com/julienschmidt/httprouter"
	handlers "github.com/veliancreate/books-api/internal/handlers/book"
	"github.com/veliancreate/books-api/internal/middleware"
	"github.com/veliancreate/books-api/internal/store"
)

func getRouter() *httprouter.Router {
	router := httprouter.New()

	auth := middleware.NewAuth()

	logger := middleware.NewLogger()

	cors := middleware.NewSecurity()

	mwareStack := []httprouter.Handle{
		logger.Init,
		auth.Authenticate,
		cors.AddHeaders,
	}

	store := store.NewInMemoryBookStore()

	booksHandler := handlers.NewBookHandler(store, logger)

	booksStack := middleware.NewMiddlewareStackHandler(
		append(mwareStack, booksHandler.ListBooks),
	)

	router.GET("/books", booksStack.Handle)

	return router
}
