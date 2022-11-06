package app_controller

import (
	"agedito/udemy/rest_api_jwt/internal/application/services/utils"
	"agedito/udemy/rest_api_jwt/internal/application/use_cases"
	"agedito/udemy/rest_api_jwt/internal/domain"
	"agedito/udemy/rest_api_jwt/internal/platform/repository"
	"encoding/json"
	"errors"
	"net/http"
)

type AppController struct {
	Repo  repository.Repository
	Cases use_cases.UseCases
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

func (c *AppController) getUserFromRequest(w http.ResponseWriter, r *http.Request) (domain.User, error) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if utils.IsError(err) {
		c.responseError(w, http.StatusBadRequest, err)
		return domain.User{}, DecodeUserError
	}

	var ok bool
	ok, err = user.Validate()
	if !ok {
		c.responseError(w, http.StatusBadRequest, err)
		return domain.User{}, err
	}

	return user, nil
}

func New(repo repository.Repository) AppController {
	return AppController{Repo: repo, Cases: use_cases.New(repo)}
}

func (c *AppController) getEmailFromTokenRequest(_ http.ResponseWriter, r *http.Request) (string, error) {
	token, err := c.getTokenFromRequest(r)
	if utils.IsError(err) {
		return "", DecodeUserError
	}

	email, emailErr := token.GetEmail()
	if utils.IsError(emailErr) {
		return "", DecodeUserError
	}
	return email, nil
}
