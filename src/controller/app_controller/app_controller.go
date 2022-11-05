package app_controller

import (
	"agedito/udemy/rest_api_jwt/models"
	"agedito/udemy/rest_api_jwt/repository"
	"agedito/udemy/rest_api_jwt/utils"
	"encoding/json"
	"errors"
	"net/http"
)

type AppController struct {
	Repo repository.Repository
}

var DecodeUserError = errors.New("error decoding user")

func (c *AppController) responseError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(err.Error())
}

func (c *AppController) responseJson(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}

func (c *AppController) getUserFromRequest(w http.ResponseWriter, r *http.Request) (models.User, error) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if utils.AssertError(err) {
		c.responseError(w, http.StatusBadRequest, err)
		return models.User{}, DecodeUserError
	}

	var ok bool
	ok, err = user.Validate()
	if !ok {
		c.responseError(w, http.StatusBadRequest, err)
		return models.User{}, err
	}

	return user, nil
}

func New(repo repository.Repository) AppController {
	return AppController{Repo: repo}
}
