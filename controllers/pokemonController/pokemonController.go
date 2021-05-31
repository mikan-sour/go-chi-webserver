package pokemonController

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jedzeins/go-chi-webserver/domains"
	"github.com/jedzeins/go-chi-webserver/services/pokemonService"
)

func GetPikachu(w http.ResponseWriter, r *http.Request) {
	pikachu := pokemonService.GetPikachu()

	res, err := json.Marshal(pikachu)
	if err != nil {
		fmt.Printf("error in marshall: %s", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllPokemon(w http.ResponseWriter, r *http.Request) {

	pokemon := pokemonService.GetAllPokemon()

	result := domains.PokemonResponseMany{
		StatusCode: http.StatusOK,
		Pokemon:    pokemon,
	}

	res, err := json.Marshal(&result)
	if err != nil {
		fmt.Printf("error in marshall: %s", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
