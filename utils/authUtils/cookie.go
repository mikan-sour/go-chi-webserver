package authUtils

import (
	"net/http"

	"github.com/jedzeins/go-chi-webserver/services/authService"
)

func ReturnAccessTokenCookie(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:     "accessToken",
		Value:    token,
		Expires:  authService.CookieExpiresTime.Add(1000000),
		Secure:   false,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
}
