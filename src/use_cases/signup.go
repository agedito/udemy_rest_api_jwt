package use_cases

import (
	"agedito/udemy/rest_api_jwt/internal/application/services/password"
	"agedito/udemy/rest_api_jwt/internal/application/services/utils"
	"agedito/udemy/rest_api_jwt/internal/domain"
	"errors"
)

var UserAlreadyExitsError = errors.New("user already exists")

func (cases *UseCases) SignUp(user domain.User) (bool, error) {
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
