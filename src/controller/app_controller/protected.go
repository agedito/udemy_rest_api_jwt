package app_controller

import (
	"fmt"
	"net/http"
)

func (_ *AppController) Protected(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("Protected")
}
