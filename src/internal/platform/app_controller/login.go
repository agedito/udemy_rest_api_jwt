package app_controller

import (
	"agedito/udemy/rest_api_jwt/internal/application/services/utils"
	"errors"
	"net/http"
)

var InvalidLoginError = errors.New("invalid email or password")

func (c *AppController) Login(w http.ResponseWriter, r *http.Request) {
	requestUser, requestErr := c.getUserFromRequest(w, r)
	if utils.IsError(requestErr) {
		c.responseError(w, http.StatusConflict, InvalidLoginError)
		return
	}

	loginToken, loginErr := c.Cases.Login(requestUser.Email, requestUser.Password)
	if utils.IsError(loginErr) {
		c.responseError(w, http.StatusConflict, InvalidLoginError)
		return
	}

	c.responseJson(w, http.StatusOK, loginToken)
}
