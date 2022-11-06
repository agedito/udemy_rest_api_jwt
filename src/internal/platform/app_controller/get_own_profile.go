package app_controller

import (
	"agedito/udemy/rest_api_jwt/internal/application/services/utils"
	"net/http"
)

func (c *AppController) GetOwnProfile(w http.ResponseWriter, r *http.Request) {
	email, err := c.getEmailFromTokenRequest(w, r)
	email, err = c.Cases.GetOwnProfile(email)
	if utils.IsError(err) {
		c.responseError(w, http.StatusUnauthorized, err)
	}

	data := make(map[string]string)
	data["email"] = email
	c.responseJson(w, http.StatusOK, data)
}
