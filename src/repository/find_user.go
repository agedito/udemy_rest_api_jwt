package repository

import "os/user"

func (r *Repository) FindUser(email string) (user.User, bool, error) {
	return user.User{}, true, nil
	//row := db.QueryRow("select * from users where email=$1", user.Email)
}
