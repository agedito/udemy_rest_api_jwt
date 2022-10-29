package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
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
	var finalError Error
	_ = json.NewDecoder(r.Body).Decode(&user)
	spew.Dump("User", user)

	if user.Email == "" {
		finalError.Message = "email is missing"
		respondWithError(w, http.StatusBadRequest, finalError)
		return
	}

	if user.Password == "" {
		finalError.Message = "password is missing"
		respondWithError(w, http.StatusBadRequest, finalError)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(hash)

	stmt := "insert into users (email, password) values($1, $2) RETURNING id;"
	err = db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		finalError.Message = "Server error."
		respondWithError(w, http.StatusInternalServerError, finalError)
		return
	}

	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	responseJSON(w, user)
}

func login(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	token, err := GenerateToken(user)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(token)
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

func responseJSON(w http.ResponseWriter, data interface{}) {
	_ = json.NewEncoder(w).Encode(data)
}

func GenerateToken(user User) (string, error) {
	var err error
	secret := "secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}
