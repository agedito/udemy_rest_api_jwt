package app_controller

import (
	"fmt"
	"net/http"
)

func (_ *AppController) SignUp(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("SignUp")
}
