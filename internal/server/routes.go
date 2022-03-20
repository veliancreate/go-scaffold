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

	bookListStack := middleware.NewMiddlewareStackHandler(
		append(mwareStack, booksHandler.ListBooks),
	)

	bookCreateStack := middleware.NewMiddlewareStackHandler(
		append(mwareStack, booksHandler.Create),
	)

	bookUpdateStack := middleware.NewMiddlewareStackHandler(
		append(mwareStack, booksHandler.Update),
	)

	bookDeleteStack := middleware.NewMiddlewareStackHandler(
		append(mwareStack, booksHandler.Delete),
	)

	router.GET("/books", bookListStack.Handle)

	router.POST("/books", bookCreateStack.Handle)

	router.PATCH("/books/:id", bookUpdateStack.Handle)

	router.DELETE("/books/:id", bookDeleteStack.Handle)

	return router
}
