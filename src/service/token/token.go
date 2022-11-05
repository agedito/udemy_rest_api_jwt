package token

import (
	"agedito/udemy/rest_api_jwt/models"
	"agedito/udemy/rest_api_jwt/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var GenerationTokenError = errors.New("error generating token")

const secret string = "secret"

type Token struct {
	Token string
}

func NewFromId(tokenId string) (Token, error) {
	_, err := getJwtToken(tokenId)
	if utils.AssertError(err) {
		return Token{}, GenerationTokenError
	}
	return Token{tokenId}, nil
}

func NewFromUser(user models.User) (Token, error) {
	// TODO: environment management
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

func getJwtToken(tokenId string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(tokenId, parseTokenMethod)
	if utils.AssertError(err) {
		return &jwt.Token{}, GenerationTokenError
	}
	if !jwtToken.Valid {
		return &jwt.Token{}, GenerationTokenError
	}
	return jwtToken, nil
}

func parseTokenMethod(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, GenerationTokenError
	}
	return []byte(secret), nil
}
