package bootstrap

import (
	"agedito/udemy/rest_api_jwt/internal/application/services/utils"
	"agedito/udemy/rest_api_jwt/internal/platform/app_controller"
	"agedito/udemy/rest_api_jwt/internal/platform/repository"
	"agedito/udemy/rest_api_jwt/internal/platform/server"
	"errors"
	"fmt"
	"log"
)

const (
	address = ":8000"
	url     = "postgres://qyiuphex:PpgU6eSADADi89kMLqLMg9Xu4SZmnwgj@lucky.db.elephantsql.com/qyiuphex"
)

var RepoError = errors.New("error creating repository")

func Run() error {

	s := server.New(address)
	repo, err := repository.New(url)
	if utils.IsError(err) {
		log.Fatal(err)
		return RepoError
	}
	c := app_controller.New(repo)

	fmt.Println("Listen at :8000...")
	//LEARN pointer parameters, when is needed?
	return s.Run(&c)
}
