package testingDogs

// avoid circular dependency so created a new package ^^

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/jedzeins/go-chi-webserver/controllers/dogController"
	"github.com/jedzeins/go-chi-webserver/controllers/marco"
	"github.com/jedzeins/go-chi-webserver/domains"
	"github.com/stretchr/testify/assert"
)

var (
	r  = chi.NewRouter()
	rr = httptest.NewRecorder()
)

func setup() {
	r.Get("/marco", marco.Marco)
	r.Get("/dogs", dogController.GetDogs)
	r.Get("/dogs/{index}", dogController.GetOneDog)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestMarcoPolo(t *testing.T) {

	req, err := http.NewRequest("GET", "/marco", nil)
	r.ServeHTTP(rr, req)

	assert.NotNil(t, req)
	assert.Nil(t, err)

	stringVal := string(rr.Body.Bytes())

	assert.EqualValues(t, "polo", stringVal)

}

func TestDogsResponseHasTwoDogs(t *testing.T) {
	req, err := http.NewRequest("GET", "/dogs", nil)
	r.ServeHTTP(rr, req)

	assert.NotNil(t, req)
	assert.Nil(t, err)
	assert.EqualValues(t, 200, rr.Result().StatusCode)

	result := make([]domains.Dog, 2)
	err = json.Unmarshal(rr.Body.Bytes(), &result)
	if err != nil {
		t.Errorf("error in unmarshalling : %v", err)
	}

	assert.EqualValues(t, "Zeus", result[0].Name)
	assert.EqualValues(t, 28, result[0].Owner.Age)
}

func TestDogsOneValNotNumber(t *testing.T) {
	req, _ := http.NewRequest("GET", "/dogs/a", nil)
	r.ServeHTTP(rr, req)

	assert.NotNil(t, rr.Body)
	assert.EqualValues(t, http.StatusBadRequest, rr.Result().StatusCode)

	result := domains.DogError{}
	err := json.Unmarshal(rr.Body.Bytes(), &result)
	if err != nil {
		t.Errorf("error in unmarshalling : %v", err)
	}

	assert.EqualValues(t, "The index needs to be an integer", result.Message)
	assert.EqualValues(t, http.StatusBadRequest, result.StatusCode)
}

func TestDogsOneValNotOneOrTwo(t *testing.T) {
	req, _ := http.NewRequest("GET", "/dogs/3", nil)
	r.ServeHTTP(rr, req)

	assert.NotNil(t, rr.Body)
	assert.EqualValues(t, http.StatusBadRequest, rr.Result().StatusCode)

	result := domains.DogError{}
	err := json.Unmarshal(rr.Body.Bytes(), &result)
	if err != nil {
		t.Errorf("error in unmarshalling : %v", err)
	}

	assert.EqualValues(t, "The index needs to be 1 or 2", result.Message)
	assert.EqualValues(t, http.StatusBadRequest, result.StatusCode)
}

func TestDogsOneSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/dogs/1", nil)
	r.ServeHTTP(rr, req)

	assert.NotNil(t, req)
	assert.Nil(t, err)
	assert.EqualValues(t, 200, rr.Result().StatusCode)

	result := domains.DogResponseOne{}
	err = json.Unmarshal(rr.Body.Bytes(), &result)
	if err != nil {
		t.Errorf("error in unmarshalling : %v", err)
	}

	assert.EqualValues(t, "OK", result.Status)
	assert.EqualValues(t, http.StatusOK, result.StatusCode)
	assert.EqualValues(t, "Zeus", result.Dog.Name)
	assert.EqualValues(t, 28, result.Dog.Owner.Age)
}
