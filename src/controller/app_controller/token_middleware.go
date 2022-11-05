package app_controller

import (
	"net/http"
)

func (c *AppController) TokenMiddleware(callback http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		callback.ServeHTTP(w, r)
	}
}
