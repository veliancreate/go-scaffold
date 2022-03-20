package middleware

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Logger struct {
	method string
	path   string
}

func NewLogger() *Logger {
	return &Logger{
		method: "NONE",
		path:   "NONE",
	}
}

func (l *Logger) Init(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	l.method = r.Method
	l.path = r.URL.EscapedPath()
}

func (l *Logger) Info(message string) {
	log.Printf("METHOD: %s; PATH: %s; MESSAGE: %s", l.method, l.path, message)
}

func (l *Logger) Error(message string) {
	log.Printf("ERROR: METHOD: %s; PATH: %s; MESSAGE: %s", l.method, l.path, message)
}
