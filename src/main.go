package main

import (
	"agedito/udemy/rest_api_jwt/bootstrap"
	"agedito/udemy/rest_api_jwt/utils"
	"log"
)

// FEATURE Manage environments
// FEATURE Create postman and .http files
// FEATURE Create tests
// FEATURE Error management

func main() {
	err := bootstrap.Run()
	if utils.IsError(err) {
		log.Fatal(err)
		return
	}
}
