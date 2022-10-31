package bootstrap

import (
	"agedito/udemy/rest_api_jwt/controller/app_controller"
	"agedito/udemy/rest_api_jwt/server"
	"fmt"
)

func Run() error {
	s := server.New(":8000")
	c := app_controller.AppController{}

	fmt.Println("Listen at :8000...")
	//LEARN pointer parameters, when is needed?
	return s.Run(&c)
}
