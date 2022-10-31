package bootstrap

import (
	"agedito/udemy/rest_api_jwt/controller/app_controller"
	"agedito/udemy/rest_api_jwt/repository"
	"agedito/udemy/rest_api_jwt/server"
	"errors"
	"fmt"
	"log"
)

const (
	address = ":8000"
	url     = "postgresql://qyiuphex:qyiuphex@lucky.db.elephantsql.com:5432/qyiuphex"
)

var RepoError = errors.New("error creating repository")

func Run() error {

	s := server.New(address)
	repo, err := repository.New(url)
	if err != nil {
		log.Fatal(err)
		return RepoError
	}
	c := app_controller.New(repo)

	fmt.Println("Listen at :8000...")
	//LEARN pointer parameters, when is needed?
	return s.Run(&c)
}
