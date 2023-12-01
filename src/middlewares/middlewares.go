package middlewares

import (
	"api/src/authentication"
	"api/src/response"
	"fmt"
	"net/http"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := authentication.ValidateToken(r); erro != nil {
			response.Error(w, http.StatusUnauthorized, erro)
			return
		}

		next(w, r)
	}
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n %s %s %s\n", r.Method, r.RequestURI, r.Host)

		next(w, r)
	}
}
