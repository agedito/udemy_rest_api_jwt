package token

import (
	"agedito/udemy/rest_api_jwt/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var GenerationTokenError = errors.New("error generating token")

type Token struct {
	Token string
}

func New(user models.User) (Token, error) {
	// TODO: environment management
	secret := "secret"

	claims := jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenId, err := token.SignedString([]byte(secret))

	if err != nil {
		return Token{}, GenerationTokenError
	}

	return Token{Token: tokenId}, nil
}
