package repository

import (
	"agedito/udemy/rest_api_jwt/utils"
	"database/sql"
	"github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func New(url string) (Repository, error) {
	repo := Repository{}
	err := repo.ConnectDB(url)
	if utils.IsError(err) {
		return repo, err
	}

	return repo, nil
}

func (r *Repository) ConnectDB(url string) error {
	pgUrl, err := pq.ParseURL(url)
	if utils.IsError(err) {
		return err
	}

	var db *sql.DB
	db, err = sql.Open("postgres", pgUrl)
	if utils.IsError(err) {
		return err
	}

	err = db.Ping()
	if utils.IsError(err) {
		return err
	}

	r.db = db
	return nil
}
