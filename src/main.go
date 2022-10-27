package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
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
	var user User
	var err Error
	_ = json.NewDecoder(r.Body).Decode(&user)
	spew.Dump("User", user)

	if user.Email == "" {
		err.Message = "email is missing"
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	if user.Password == "" {
		err.Message = "password is missing"
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
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

func respondWithError(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(error)
}
