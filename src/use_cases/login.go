package use_cases

import (
	"agedito/udemy/rest_api_jwt/service/password"
	"agedito/udemy/rest_api_jwt/service/token"
	"agedito/udemy/rest_api_jwt/utils"
	"errors"
)

var InvalidEmailError = errors.New("invalid email")
var InvalidPasswordError = errors.New("invalid password")
var GeneratingTokenError = errors.New("error generating token")

func (cases *UseCases) Login(userEmail string, userPassword string) (token.Token, error) {
	repoUser, exists, _ := cases.Repo.FindUser(userEmail)
	if !exists {
		return token.Token{}, InvalidEmailError
	}

	hashedPassword := repoUser.Password
	match, passwordErr := password.CheckPassword(userPassword, hashedPassword)
	if !match || utils.IsError(passwordErr) {
		return token.Token{}, InvalidPasswordError
	}

	loginToken, tokenErr := token.NewFromUser(repoUser)
	if utils.IsError(tokenErr) {
		return token.Token{}, GeneratingTokenError
	}

	return loginToken, nil
}
