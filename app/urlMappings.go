package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/jedzeins/go-chi-webserver/controllers/authController"
	"github.com/jedzeins/go-chi-webserver/controllers/dogController"
	"github.com/jedzeins/go-chi-webserver/controllers/marco"
	"github.com/jedzeins/go-chi-webserver/controllers/pokemonController"
	"github.com/jedzeins/go-chi-webserver/middleware/auth_middleware"
)

func DoURLMappings(router chi.Router) {
	router.Get("/marco", marco.Marco)
	router.Get("/dogs", dogController.GetDogs)
	router.Get("/dogs/{index}", dogController.GetOneDog)
	router.Get("/pikachu", pokemonController.GetPikachu)
	router.Get("/pokemon", pokemonController.GetAllPokemon)

	router.Group(func(router chi.Router) {
		router.Use(auth_middleware.AuthMiddleware) // implement the middleware...
		router.Get("/midwareTest", marco.Marco)
		router.Get("/authParse", authController.ParseJWT)

	})
	router.Get("/authToken", authController.NewJWT)

}
