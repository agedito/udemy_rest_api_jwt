package use_cases

import (
	"agedito/udemy/rest_api_jwt/models"
	"agedito/udemy/rest_api_jwt/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var UserAlreadyExitsError = errors.New("user already exists")
var GeneratingPasswordError = errors.New("error generating password")

func (cases *UseCases) SignUp(user models.User) (bool, error) {
	_, exists, _ := cases.Repo.FindUser(user.Email)
	if exists {
		return false, UserAlreadyExitsError
	}

	// TODO: password service
	hash, passwordErr := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if utils.IsError(passwordErr) {
		return false, GeneratingPasswordError
	}

	user.Password = string(hash)
	created, repoCreationErr := cases.Repo.CreateUser(user)
	if !created {
		return false, repoCreationErr
	}

	return true, nil
}
