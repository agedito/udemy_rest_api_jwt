package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	router *mux.Router
}

func New() Server {
	server := Server{}

	server.router = mux.NewRouter()
	server.createEndpoints()

	return server
}

func (server *Server) Run() error {
	return http.ListenAndServe(":8000", server.router)
}

func (server *Server) createEndpoints() {
	server.router.HandleFunc("/ping", server.ping)
}

func (server *Server) ping(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("Pong")
}
