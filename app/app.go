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
	doURLMappings(router)

	fmt.Printf("serving on port %s\n", port)
	http.ListenAndServe(port, router)
}
