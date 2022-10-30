package main

import (
	"agedito/udemy/rest_api_jwt/driver"
	"agedito/udemy/rest_api_jwt/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strings"
)

var db *sql.DB

func init() {
	_ = gotenv.Load()
}

func main() {
	db = driver.ConnectDB()

	router := mux.NewRouter()
	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleWare(ProtectedEndPoint)).Methods("POST")

	log.Println("Listen on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var finalError models.Error
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
	var user models.User
	var resultJwt models.JWT
	var err models.Error

	_ = json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		err.Message = "Email is missing."
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	if user.Password == "" {
		err.Message = "Password is missing."
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	password := user.Password

	row := db.QueryRow("select * from users where email=$1", user.Email)
	sqlErr := row.Scan(&user.ID, &user.Email, &user.Password)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			err.Message = "The user does not exist"
			respondWithError(w, http.StatusBadRequest, err)
			return
		} else {
			log.Fatal(sqlErr)
		}
	}

	hashedPassword := user.Password

	sqlErr = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if sqlErr != nil {
		err.Message = "Invalid Password"
		respondWithError(w, http.StatusUnauthorized, err)
		return
	}

	token, sqlErr := GenerateToken(user)

	if sqlErr != nil {
		log.Fatal(sqlErr)
	}

	w.WriteHeader(http.StatusOK)
	resultJwt.Token = token

	responseJSON(w, resultJwt)
}

func ProtectedEndPoint(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("Protected endpoint invoked")
}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var errorObject models.Error
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}

				return []byte(os.Getenv("secret")), nil
			})

			if err != nil {
				errorObject.Message = err.Error()
				respondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				errorObject.Message = err.Error()
				respondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}
		} else {
			errorObject.Message = "Invalid token."
			respondWithError(w, http.StatusUnauthorized, errorObject)
			return
		}
	}
}

func respondWithError(w http.ResponseWriter, status int, error models.Error) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(error)
}

func responseJSON(w http.ResponseWriter, data interface{}) {
	_ = json.NewEncoder(w).Encode(data)
}

func GenerateToken(user models.User) (string, error) {
	var err error

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("secret")))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}
