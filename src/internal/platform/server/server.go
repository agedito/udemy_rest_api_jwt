package server

import (
	"agedito/udemy/rest_api_jwt/internal/platform/controller"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	address string
	router  *mux.Router
}

func New(address string) Server {
	server := Server{}

	server.address = address
	server.router = mux.NewRouter()

	return server
}

func (server *Server) Run(controller controller.Controller) error {
	server.createEndpoints(controller)
	return http.ListenAndServe(server.address, server.router)
}

func (server *Server) createEndpoints(controller controller.Controller) {
	//FEATURE Management 404 error, try to access to a no valid endpoint
	server.router.HandleFunc("/ping", controller.Ping)
	server.router.HandleFunc("/signup", controller.SignUp)
	server.router.HandleFunc("/login", controller.Login)
	server.router.HandleFunc("/profile", controller.TokenMiddleware(controller.GetOwnProfile))
}
