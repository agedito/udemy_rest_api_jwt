package app_controller

import (
	"fmt"
	"net/http"
)

func (_ *AppController) Login(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("Login")
}
