package app_controller

import (
	"agedito/udemy/rest_api_jwt/use_cases"
	"agedito/udemy/rest_api_jwt/utils"
	"errors"
	"fmt"
	"net/http"
)

var InternalError = errors.New("internal error")

func (c *AppController) Ping(w http.ResponseWriter, _ *http.Request) {
	message, pingErr := use_cases.Ping()
	if utils.IsError(pingErr) {
		c.responseError(w, http.StatusInternalServerError, InternalError)
	}

	_, err := fmt.Fprintf(w, message)
	if utils.IsError(err) {
		c.responseError(w, http.StatusInternalServerError, InternalError)
	}
}
