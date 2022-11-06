package app_controller

import (
	"agedito/udemy/rest_api_jwt/internal/application/services/utils"
	"net/http"
)

func (c *AppController) SignUp(w http.ResponseWriter, r *http.Request) {
	user, err := c.getUserFromRequest(w, r)
	if utils.IsError(err) {
		return
	}

	created, signupErr := c.Cases.SignUp(user)
	if !created {
		c.responseError(w, http.StatusInternalServerError, signupErr)
		return
	}

	c.responseJson(w, http.StatusOK, user)
}
