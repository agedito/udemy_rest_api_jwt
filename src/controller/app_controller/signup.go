package app_controller

import (
	"agedito/udemy/rest_api_jwt/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var UserAlreadyExitsError = errors.New("user already exists")

func (c *AppController) SignUp(w http.ResponseWriter, r *http.Request) {
	user, err := c.getUserFromRequest(w, r)
	if utils.AssertError(err) {
		return
	}

	_, exists, _ := c.Repo.FindUser(user.Email)
	if exists {
		// TODO: propagate result error
		c.responseError(w, http.StatusConflict, UserAlreadyExitsError)
		return
	}

	// TODO: password service
	hash, passwordErr := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if passwordErr != nil {
		log.Fatal(err)
	}
	user.Password = string(hash)

	created, repoCreationErr := c.Repo.CreateUser(user)
	if !created {
		c.responseError(w, http.StatusInternalServerError, repoCreationErr)
		return
	}

	c.responseJson(w, http.StatusOK, user)
}
