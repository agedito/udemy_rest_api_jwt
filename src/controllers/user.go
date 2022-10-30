package controllers

import (
	"agedito/udemy/rest_api_jwt/models"
	userRepository "agedito/udemy/rest_api_jwt/repository/user"
	"agedito/udemy/rest_api_jwt/utils"
	"database/sql"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type Controller struct{}

func (c Controller) SignUp(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var finalError models.Error
		_ = json.NewDecoder(r.Body).Decode(&user)
		spew.Dump("User", user)

		if user.Email == "" {
			finalError.Message = "email is missing"
			utils.RespondWithError(w, http.StatusBadRequest, finalError)
			return
		}

		if user.Password == "" {
			finalError.Message = "password is missing"
			utils.RespondWithError(w, http.StatusBadRequest, finalError)
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
			utils.RespondWithError(w, http.StatusInternalServerError, finalError)
			return
		}

		user.Password = ""
		w.Header().Set("Content-Type", "application/json")
		utils.ResponseJSON(w, user)
	}
}

func (c Controller) Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		userRepo := userRepository.UserRepository{}
		user, loginErr := userRepo.Login(db, user)
		if loginErr != nil {
			log.Fatal(loginErr)
		}
		hashedPassword := user.Password

		sqlErr := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

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
}
