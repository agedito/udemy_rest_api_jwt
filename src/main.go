package main

import "agedito/udemy/rest_api_jwt/bootstrap"

func main() {
	err := bootstrap.Run()
	if err != nil {
		return
	}
}
