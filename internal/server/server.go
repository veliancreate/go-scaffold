package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/veliancreate/books-api/internal/config"
)

func Start() {
	router := getRouter()

	config := config.NewConfig()

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("127.0.0.1:%s", config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("starting server on port %s ...", config.Port)

	log.Fatal(srv.ListenAndServe())
}
