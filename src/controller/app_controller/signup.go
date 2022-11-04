package app_controller

import (
	"agedito/udemy/rest_api_jwt/models"
	"agedito/udemy/rest_api_jwt/utils"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var UserAlreadyExits = errors.New("user already exists")

func (c *AppController) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if utils.AssertError(err) {
		c.responseError(w, http.StatusBadRequest, err)
		return
	}

	var ok bool
	ok, err = user.Validate()
	if !ok {
		c.responseError(w, http.StatusBadRequest, err)
		return
	}

	_, exists, _ := c.Repo.FindUser(user.Email)
	if exists {
		// TODO: propagate result error
		c.responseError(w, http.StatusConflict, UserAlreadyExits)
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

	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	c.ResponseJson(w, user)
}
