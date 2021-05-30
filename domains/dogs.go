package domains

type Dog struct {
	Name       string   `json:"name"`
	Breed      string   `json:"breed"`
	LikesToEat []string `json:"likesToEat"`
	Owner      Owner    `json:"owner"`
}

type Owner struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type DogResponseMany struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Dogs       []Dog  `json:"dogs"`
}

type DogResponseOne struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Dog        Dog    `json:"dog"`
}

type DogError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
