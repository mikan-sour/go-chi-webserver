package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/jedzeins/go-chi-webserver/controllers/dogController"
	"github.com/jedzeins/go-chi-webserver/controllers/marco"
)

func DoURLMappings(router *chi.Mux) {
	router.Get("/marco", marco.Marco)
	router.Get("/dogs", dogController.GetDogs)
	router.Get("/dogs/{index}", dogController.GetOneDog)
}
