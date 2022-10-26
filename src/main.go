package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"log"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}

var db *sql.DB

func main() {
	pgUrl, err := pq.ParseURL("postgres://qyiuphex:PpgU6eSADADi89kMLqLMg9Xu4SZmnwgj@lucky.db.elephantsql.com/qyiuphex")
	if err != nil {
		log.Fatal(err)
		return
	}

	db, err = sql.Open("postgres", pgUrl)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}

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
