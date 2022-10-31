package bootstrap

import (
	"agedito/udemy/rest_api_jwt/server"
)

func Run() error {
	s := server.New()
	return s.Run()
}
