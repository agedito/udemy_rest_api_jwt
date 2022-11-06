package use_cases

import (
	"agedito/udemy/rest_api_jwt/models"
	"agedito/udemy/rest_api_jwt/service/password"
	"agedito/udemy/rest_api_jwt/utils"
	"errors"
)

var UserAlreadyExitsError = errors.New("user already exists")

func (cases *UseCases) SignUp(user models.User) (bool, error) {
	_, exists, _ := cases.Repo.FindUser(user.Email)
	if exists {
		return false, UserAlreadyExitsError
	}

	hashPassword, passwordErr := password.GeneratePassword(user.Password)
	if utils.IsError(passwordErr) {
		return false, passwordErr
	}

	user.Password = hashPassword
	created, repoCreationErr := cases.Repo.CreateUser(user)
	if !created {
		return false, repoCreationErr
	}

	return true, nil
}
