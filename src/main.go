package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleWare(ProtectedEndPoint)).Methods("POST")

	log.Println("Listen on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SignUp invoked")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login invoked")
}

func ProtectedEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Protected endpoint invoked")
}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	callback := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Token Verified!!")
		next.ServeHTTP(w, r)
	}
	return callback
}
