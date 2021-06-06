package authController

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/jedzeins/go-chi-webserver/domains"
	"github.com/jedzeins/go-chi-webserver/services/authService"
	"github.com/jedzeins/go-chi-webserver/utils/authUtils"
)

func NewJWT(w http.ResponseWriter, r *http.Request) {

	var mockDBResponse = domains.UserDBResponse{
		UserName:     "jedzeins",
		TokenVersion: 1,
	}

	claims, err := authService.CreateClaims(mockDBResponse)
	if err != nil {
		fmt.Printf("error in the create claims func: %w", err)
	}

	token, err := authService.CreateToken(claims)

	if err != nil {
		errString := fmt.Sprintf("error running service: %s", err)
		w.Write([]byte(errString))
		return
	}

	authUtils.ReturnAccessTokenCookie(w, token)

	w.Write([]byte(token))
}

func ParseJWT(w http.ResponseWriter, r *http.Request) {

	tokenUnformatted := r.Header.Get("Authorization")

	if tokenUnformatted == "" {
		err := errors.New("no authorization token provided")
		errRes := domains.AuthError{
			Error:      fmt.Sprintf("Error: %s", err),
			StatusCode: http.StatusUnauthorized,
		}

		res, err := json.Marshal(errRes)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(401)
		w.Write(res)
		return
	}

	token := strings.Split(r.Header.Get("Authorization"), " ")[1]

	parsed, err := authService.ParseToken(token)
	if err != nil {
		fmt.Printf("error in marshall: %s", err)
	}

	res, err := json.Marshal(parsed)
	if err != nil {
		fmt.Printf("error in marshall: %s", err)
	}
	w.WriteHeader(200)
	w.Write(res)
}
