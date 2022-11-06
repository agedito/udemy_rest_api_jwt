package token

import (
	"agedito/udemy/rest_api_jwt/models"
	"agedito/udemy/rest_api_jwt/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var GenerationTokenError = errors.New("error generating token")
var InvalidTokenError = errors.New("invalid token")

const secret string = "secret"

type Token struct {
	Token string
}

type UserClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
	Iss   string `json:"iss"`
}

func NewFromId(tokenId string) (Token, error) {
	_, err := getJwtToken(tokenId)
	if utils.IsError(err) {
		return Token{}, GenerationTokenError
	}
	return Token{tokenId}, nil
}

func NewFromUser(user models.User) (Token, error) {
	claims := UserClaims{Email: user.Email, Iss: "course"}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenId, err := token.SignedString([]byte(secret))

	if utils.IsError(err) {
		return Token{}, GenerationTokenError
	}

	return Token{Token: tokenId}, nil
}

func getJwtToken(tokenId string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(tokenId, parseTokenMethod)
	if utils.IsError(err) {
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

func (t *Token) GetEmail() (string, error) {
	claims := UserClaims{}
	_, err := jwt.ParseWithClaims(t.Token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if utils.IsError(err) {
		return "", InvalidTokenError
	}
	return claims.Email, nil
}
