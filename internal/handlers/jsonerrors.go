package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ErrorLogger interface {
	Error(message string)
}

// influenced by https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
func HandleJSONParsingError(err error, w http.ResponseWriter, errorLogger ErrorLogger) {

	var syntaxError *json.SyntaxError

	var unmarshalTypeError *json.UnmarshalTypeError

	switch {
	case errors.As(err, &syntaxError):
		http.Error(w, "json syntax error", http.StatusBadRequest)
		return
	case errors.Is(err, io.ErrUnexpectedEOF):
		msg := "request body contains badly-formed JSON"
		http.Error(w, msg, http.StatusBadRequest)
		return
	case errors.As(err, &unmarshalTypeError):
		msg := "unmarshal type error"
		http.Error(w, msg, http.StatusBadRequest)
		return
	case strings.HasPrefix(err.Error(), "json: unknown field "):
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
		msg := fmt.Sprintf("request body contains unknown field %s", fieldName)
		http.Error(w, msg, http.StatusBadRequest)
		return
	case errors.Is(err, io.EOF):
		msg := "Request body must not be empty"
		http.Error(w, msg, http.StatusBadRequest)
		return
	case err.Error() == "http: request body too large":
		msg := "request body too large"
		http.Error(w, msg, http.StatusRequestEntityTooLarge)
		return
	default:
		errorLogger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
