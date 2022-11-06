package models

import (
	"agedito/udemy/rest_api_jwt/utils"
	"errors"
	"net/mail"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var EmptyEmailError = errors.New("email is missing")
var InvalidEmailError = errors.New("invalid email")
var EmptyPasswordsError = errors.New("password is missing")

func (user *User) Validate() (bool, error) {
	if user.Email == "" {
		return false, EmptyEmailError
	}

	if user.Password == "" {
		return false, EmptyPasswordsError
	}

	_, err := mail.ParseAddress(user.Email)
	if utils.IsError(err) {
		return false, InvalidEmailError
	}

	if !strings.Contains(user.Email, ".") {
		return false, InvalidEmailError
	}

	return true, nil
}
