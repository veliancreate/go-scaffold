package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Security struct{}

func NewSecurity() *Security {
	return &Security{}
}

func (c *Security) AddHeaders(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
