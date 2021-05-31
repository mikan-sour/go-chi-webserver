package pokemonProviders

import (
	"github.com/jedzeins/go-chi-webserver/domains"
)

var client domains.DB

func init() {
	client.PostgresConnect()
}
