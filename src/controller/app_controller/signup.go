package app_controller

import (
	"agedito/udemy/rest_api_jwt/models"
	"agedito/udemy/rest_api_jwt/utils"
	"encoding/json"
	"net/http"
)

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

	c.ResponseJson(w, user)

}
