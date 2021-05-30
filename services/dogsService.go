package services

import (
	"net/http"

	"github.com/jedzeins/go-chi-webserver/domains"
)

var dogs = []domains.Dog{
	{Name: "Zeus", Breed: "Great Pyranees", LikesToEat: []string{"biscuts", "kibbles", "bacon"}, Owner: domains.Owner{Name: "Jed", Age: 28}},
	{Name: "Zoe", Breed: "Schnoodle", LikesToEat: []string{"Dog food", "One fried egg"}, Owner: domains.Owner{Name: "Molly", Age: 33}},
}

func GetDogs() []domains.Dog {
	return dogs
}

func GetOneDog(index int) (*domains.DogResponseOne, *domains.DogError) {
	if index > 2 || index < 1 {
		apiError := domains.DogError{
			StatusCode: http.StatusBadRequest,
			Message:    "The index needs to be 1 or 2",
		}

		return nil, &apiError
	}

	dogResponse := domains.DogResponseOne{
		Status:     "OK",
		StatusCode: http.StatusOK,
		Dog:        dogs[index-1],
	}

	return &dogResponse, nil
}
