package controllers

import (
	"fmt"
	"net/http"
)

func (c Controller) ProtectedEndPoint(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("Protected endpoint invoked")
}
