package repository

import (
	"agedito/udemy/rest_api_jwt/internal/domain"
)

func (r *Repository) CreateUser(user domain.User) (bool, error) {
	stmt := "insert into users (email, password) values($1, $2) RETURNING id;"
	err := r.db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)
	return err == nil, err
}
