package app_controller

import (
	"agedito/udemy/rest_api_jwt/service/token"
	"agedito/udemy/rest_api_jwt/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var InvalidLoginError = errors.New("invalid email or password")

// FEATURE Implement password service
func (c *AppController) Login(w http.ResponseWriter, r *http.Request) {
	requestUser, err := c.getUserFromRequest(w, r)
	if utils.IsError(err) {
		return
	}

	repoUser, exists, _ := c.Repo.FindUser(requestUser.Email)
	if !exists {
		c.responseError(w, http.StatusConflict, InvalidLoginError)
		return
	}

	hashedPassword := repoUser.Password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(requestUser.Password))
	if err != nil {
		c.responseError(w, http.StatusUnauthorized, InvalidLoginError)
		return
	}

	loginToken, tokenErr := token.NewFromUser(requestUser)
	if utils.IsError(tokenErr) {
		c.responseError(w, http.StatusInternalServerError, tokenErr)
		return
	}

	c.responseJson(w, http.StatusOK, loginToken)
}
