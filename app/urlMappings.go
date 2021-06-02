package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/jedzeins/go-chi-webserver/controllers/authController"
	"github.com/jedzeins/go-chi-webserver/controllers/dogController"
	"github.com/jedzeins/go-chi-webserver/controllers/marco"
	"github.com/jedzeins/go-chi-webserver/controllers/pokemonController"
)

func DoURLMappings(router *chi.Mux) {
	router.Get("/marco", marco.Marco)
	router.Get("/dogs", dogController.GetDogs)
	router.Get("/dogs/{index}", dogController.GetOneDog)
	router.Get("/pikachu", pokemonController.GetPikachu)
	router.Get("/pokemon", pokemonController.GetAllPokemon)
	router.Get("/authToken", authController.TestJWT)
	router.Get("/authParse", authController.ParseJWT)
}
