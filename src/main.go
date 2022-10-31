package main

import (
	"agedito/udemy/rest_api_jwt/bootstrap"
	"log"
)

//FEATURE Manage environments
//FEATURE Create postman and .http files
//FEATURE Create tests

func main() {
	err := bootstrap.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
