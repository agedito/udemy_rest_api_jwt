package password

import (
	"agedito/udemy/rest_api_jwt/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var GeneratingPasswordError = errors.New("error generating password")
var InvalidPasswordError = errors.New("invalid password")

func GeneratePassword(plainPassword string) (string, error) {
	hash, passwordErr := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if utils.IsError(passwordErr) {
		return "", GeneratingPasswordError
	}

	return string(hash), nil
}

func CheckPassword(plainPassword string, hashPassword string) (bool, error) {
	passwordErr := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))
	if utils.IsError(passwordErr) {
		return false, InvalidPasswordError
	}

	return true, nil
}
