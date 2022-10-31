package app_controller

import (
	"fmt"
	"net/http"
)

// FEATURE Implement this
func (_ *AppController) Login(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("Login")
}
