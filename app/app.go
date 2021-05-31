package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	port = ":8080"
)

func StartChiApp() {

	router := chi.NewRouter()
	DoURLMappings(router)

	fmt.Printf("serving on port %s\n", port)
	http.ListenAndServe(port, router)
}
