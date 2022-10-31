package app_controller

import (
	"agedito/udemy/rest_api_jwt/utils"
	"fmt"
	"net/http"
)

func (_ *AppController) Ping(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Pong")
	if utils.AssertError(err) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "internal error")
	}
}
