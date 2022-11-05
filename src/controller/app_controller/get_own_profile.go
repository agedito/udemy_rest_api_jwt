package app_controller

import (
	"agedito/udemy/rest_api_jwt/utils"
	"errors"
	"net/http"
)

var InvalidEmailError = errors.New("invalid error")
var NotFindUserError = errors.New("no find user")

func (c *AppController) GetOwnProfile(w http.ResponseWriter, r *http.Request) {
	email, err := c.getEmailFromTokenRequest(w, r)
	if utils.AssertError(err) {
		c.responseError(w, http.StatusUnauthorized, InvalidEmailError)
	}

	user, exists, repoErr := c.Repo.FindUser(email)
	if !exists {
		c.responseError(w, http.StatusUnauthorized, repoErr)
	}

	if !exists {
		c.responseError(w, http.StatusUnauthorized, NotFindUserError)
	}

	data := make(map[string]string)
	data["email"] = user.Email
	c.responseJson(w, http.StatusOK, data)
}
