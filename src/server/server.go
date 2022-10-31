package server

import (
	"fmt"
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
	server.createEndpoints()

	return server
}

func (server *Server) Run() error {
	return http.ListenAndServe(server.address, server.router)
}

func (server *Server) createEndpoints() {
	server.router.HandleFunc("/ping", server.ping)
}

func (server *Server) ping(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("Pong")
}
