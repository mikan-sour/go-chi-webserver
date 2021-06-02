package authController

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/jedzeins/go-chi-webserver/domains"
	"github.com/jedzeins/go-chi-webserver/services/authService"
)

func TestJWT(w http.ResponseWriter, r *http.Request) {

	claims := domains.UserClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        "jedzeins",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
		SessionId: 123,
	}
	token, err := authService.CreateToken(&claims)

	if err != nil {
		errString := fmt.Sprintf("error running service: %s", err)
		w.Write([]byte(errString))
		return
	}

	w.Write([]byte(token))
}

func ParseJWT(w http.ResponseWriter, r *http.Request) {

	claims := domains.UserClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        "jedzeins",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
		SessionId: 123,
	}
	token, err := authService.CreateToken(&claims)

	if err != nil {
		errString := fmt.Sprintf("error running service: %s", err)
		w.Write([]byte(errString))
		return
	}

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
