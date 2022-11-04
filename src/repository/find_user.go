package repository

import (
	"agedito/udemy/rest_api_jwt/models"
)

func (r *Repository) FindUser(email string) (models.User, bool, error) {
	var u models.User
	query := "select * from users where email=$1"
	row := r.db.QueryRow(query, email)
	err := row.Scan(&u.ID, &u.Email, &u.Password)
	return u, err == nil, err
}
