package dogController

import (
	"github.com/go-chi/chi/v5"
	"github.com/jedzeins/go-chi-webserver/domains"
	"github.com/jedzeins/go-chi-webserver/services"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetDogs(w http.ResponseWriter, r *http.Request) {
	dogs := services.GetDogs()

	res, err := json.Marshal(dogs)
	if err != nil {
		fmt.Printf("error in marshall: %s", err)
	}

	w.Write(res)
}

func GetOneDog(w http.ResponseWriter, r *http.Request) {

	index, errStrconv := strconv.Atoi(chi.URLParam(r, "index"))
	if errStrconv != nil {
		apiError := domains.DogError{
			StatusCode: http.StatusBadRequest,
			Message:    "The index needs to be an integer",
		}

		res, _ := json.Marshal(apiError)
		w.Write(res)
		return
	}

	dog, err := services.GetOneDog(index)
	if err != nil {
		res, _ := json.Marshal(err)
		w.Write(res)
	} else {
		res, err := json.Marshal(dog)
		if err != nil {
			fmt.Printf("error in marshall: %s", err)
		}

		w.Write(res)
	}

}
