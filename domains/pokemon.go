package domains

type Pokemon struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (Pokemon) TableName() string {
	return "pokemon"
}

type PokemonResponseMany struct {
	StatusCode int       `json:"statusCode"`
	Pokemon    []Pokemon `json:"pokemon"`
}

type PokemonResponseOne struct {
	StatusCode int     `json:"statusCode"`
	Pokemon    Pokemon `json:"pokemon"`
}

type PokemonError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
