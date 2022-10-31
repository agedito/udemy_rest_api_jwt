package app_controller

import (
	"agedito/udemy/rest_api_jwt/repository"
	"encoding/json"
	"net/http"
)

type AppController struct {
	Repo repository.Repository
}

func (c *AppController) responseError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(err.Error())
}

func (c *AppController) ResponseJson(w http.ResponseWriter, data interface{}) {
	_ = json.NewEncoder(w).Encode(data)
}

func New(repo repository.Repository) AppController {
	return AppController{Repo: repository.Repository{}}
}
