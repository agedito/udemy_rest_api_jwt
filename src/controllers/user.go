package controllers

import (
	"agedito/udemy/rest_api_jwt/models"
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
