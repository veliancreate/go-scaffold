package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MiddlewareStackHandler struct {
	stack []httprouter.Handle
}

func NewMiddlewareStackHandler(stack []httprouter.Handle) *MiddlewareStackHandler {
	return &MiddlewareStackHandler{
		stack,
	}
}

func (msh *MiddlewareStackHandler) Handle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	for i := 0; i < len(msh.stack); i++ {
		f := msh.stack[i]

		f(w, r, p)
	}
}

func (msh *MiddlewareStackHandler) Add(handler httprouter.Handle) {
	msh.stack = append(msh.stack, handler)
}
