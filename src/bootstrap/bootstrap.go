package bootstrap

import (
	"agedito/udemy/rest_api_jwt/server"
)

func Run() error {
	s := server.New(":8000")
	return s.Run()
}
