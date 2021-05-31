package pokemonService

import (
	"github.com/jedzeins/go-chi-webserver/domains"
	"github.com/jedzeins/go-chi-webserver/providers/pokemonProviders"
)

// service should handle the response

func GetPikachu() domains.Pokemon {
	return pokemonProviders.ProvidePikachu()
}

func GetAllPokemon() []domains.Pokemon {
	return pokemonProviders.ProvideAll()
}
