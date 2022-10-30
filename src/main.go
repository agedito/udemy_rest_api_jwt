package main

import (
	"agedito/udemy/rest_api_jwt/controllers"
	"agedito/udemy/rest_api_jwt/driver"
	"agedito/udemy/rest_api_jwt/models"
	"agedito/udemy/rest_api_jwt/utils"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	_ = gotenv.Load()
}

func main() {
	db = driver.ConnectDB()
	router := mux.NewRouter()

	controller := controllers.Controller{}

	router.HandleFunc("/signup", controller.SignUp(db)).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", utils.TokenVerifyMiddleWare(controller.ProtectedEndPoint)).Methods("POST")

	log.Println("Listen on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var resultJwt models.JWT
	var err models.Error

	_ = json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		err.Message = "Email is missing."
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	if user.Password == "" {
		err.Message = "Password is missing."
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	password := user.Password

	row := db.QueryRow("select * from users where email=$1", user.Email)
	sqlErr := row.Scan(&user.ID, &user.Email, &user.Password)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			err.Message = "The user does not exist"
			utils.RespondWithError(w, http.StatusBadRequest, err)
			return
		} else {
			log.Fatal(sqlErr)
		}
	}

	hashedPassword := user.Password

	sqlErr = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if sqlErr != nil {
		err.Message = "Invalid Password"
		utils.RespondWithError(w, http.StatusUnauthorized, err)
		return
	}

	token, sqlErr := utils.GenerateToken(user)

	if sqlErr != nil {
		log.Fatal(sqlErr)
	}

	w.WriteHeader(http.StatusOK)
	resultJwt.Token = token

	utils.ResponseJSON(w, resultJwt)
}
