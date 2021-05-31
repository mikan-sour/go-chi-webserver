package pokemonProviders

import (
	"github.com/jedzeins/go-chi-webserver/domains"
)

func ProvidePikachu() domains.Pokemon {
	var pikachu domains.Pokemon
	client.Database.First(&pikachu)
	return pikachu
}

func ProvideAll() []domains.Pokemon {
	var pokemon []domains.Pokemon
	client.Database.Find(&pokemon)
	return pokemon
}
