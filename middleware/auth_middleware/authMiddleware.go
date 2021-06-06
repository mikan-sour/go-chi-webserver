package auth_middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/jedzeins/go-chi-webserver/domains"
	"github.com/jedzeins/go-chi-webserver/services/authService"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println(r.Context())

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

		fmt.Printf("token: %s", parsed.Id)

		next.ServeHTTP(w, r)

	})
}
