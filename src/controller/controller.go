package controller

import "net/http"

type Controller interface {
	Ping(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	GetOwnProfile(w http.ResponseWriter, r *http.Request)

	TokenMiddleware(callback http.HandlerFunc) http.HandlerFunc
}
