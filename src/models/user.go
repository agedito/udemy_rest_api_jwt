package models

import (
	"errors"
	"net/mail"
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

	// learn check not . case
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return false, InvalidEmailError
	}

	return true, nil
}
