package use_cases

import (
	"agedito/udemy/rest_api_jwt/service/token"
	"agedito/udemy/rest_api_jwt/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var InvalidEmailError = errors.New("invalid email")
var InvalidPasswordError = errors.New("invalid password")
var GeneratingTokenError = errors.New("error generating token")

func (c *UseCases) Login(userEnail string, userPassword string) (token.Token, error) {
	repoUser, exists, _ := c.Repo.FindUser(userEnail)
	if !exists {
		return token.Token{}, InvalidEmailError
	}

	hashedPassword := repoUser.Password
	passwordErr := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword))
	if utils.IsError(passwordErr) {
		return token.Token{}, InvalidPasswordError
	}

	loginToken, tokenErr := token.NewFromUser(repoUser)
	if utils.IsError(tokenErr) {
		return token.Token{}, GeneratingTokenError
	}

	return loginToken, nil
}
