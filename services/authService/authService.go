package authService

import (
	"fmt"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/jedzeins/go-chi-webserver/domains"
)

var (
	jwtKey            = []byte("my_secret_key")
	CookieExpiresTime = time.Now().Add(time.Minute * 10)
)

func CreateClaims(dbResponse domains.UserDBResponse) (*domains.UserClaims, error) {

	claims := domains.UserClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        dbResponse.UserName,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: CookieExpiresTime.Unix(),
		},
		SessionId:    123,
		TokenVersion: dbResponse.TokenVersion,
	}

	return &claims, nil
}

func CreateToken(claims *domains.UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("error making signed token: %w", err)
	}

	return signedToken, nil

}

func ParseToken(signedToken string) (*domains.UserClaims, error) {
	claims := &domains.UserClaims{}

	t, err := jwt.ParseWithClaims(signedToken, claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("invalid signing algorithm")
		}

		return jwtKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error in parseToken: %w", err)
	}

	if !t.Valid {
		return nil, fmt.Errorf("error in parseToken - not valid : %w", err)
	}

	return t.Claims.(*domains.UserClaims), nil

}
