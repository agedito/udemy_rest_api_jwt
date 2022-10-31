package controller

import "net/http"

type Controller interface {
	Ping(w http.ResponseWriter, r *http.Request)
}
