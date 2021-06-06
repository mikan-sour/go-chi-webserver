package pokemonProviders

import (
	"fmt"

	"github.com/jedzeins/go-chi-webserver/domains"
)

var client domains.DB

func init() {
	// client.PostgresConnect()
	fmt.Println("skip DB")
}
