package utils

import (
	"agedito/udemy/rest_api_jwt/models"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"strings"
)

func RespondWithError(w http.ResponseWriter, status int, error models.Error) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(error)
}

func ResponseJSON(w http.ResponseWriter, data interface{}) {
	_ = json.NewEncoder(w).Encode(data)
}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var errorObject models.Error
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}

				return []byte(os.Getenv("secret")), nil
			})

			if err != nil {
				errorObject.Message = err.Error()
				RespondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				errorObject.Message = err.Error()
				RespondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}
		} else {
			errorObject.Message = "Invalid token."
			RespondWithError(w, http.StatusUnauthorized, errorObject)
			return
		}
	}
}

func GenerateToken(user models.User) (string, error) {
	var err error

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("secret")))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}
