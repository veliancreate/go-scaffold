package middleware

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) Authenticate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("authenticating user ...")
	log.Println("authenticated user ...")
}
