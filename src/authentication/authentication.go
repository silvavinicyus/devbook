package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.SecretKey))
}

func ValidateToken(r *http.Request) error {
	tokenString := extrairToken(r)

	token, erro := jwt.Parse(tokenString, returnValidationKey)

	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func GetUserIdFromRequest(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)

	token, erro := jwt.Parse(tokenString, returnValidationKey)

	if erro != nil {
		return 0, erro
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return userId, nil
	}

	return 0, errors.New("invalid token")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	tokenSplitted := strings.Split(token, " ")

	if len(tokenSplitted) == 2 {
		return tokenSplitted[1]
	}

	return ""
}

func returnValidationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signature method!")
	}

	return config.SecretKey, nil
}
