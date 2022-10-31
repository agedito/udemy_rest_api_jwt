package repository

import (
	"database/sql"
	"github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func New(url string) (Repository, error) {
	repo := Repository{}
	err := repo.ConnectDB(url)
	if err != nil {
		return repo, err
	}

	return repo, nil
}

func (r *Repository) ConnectDB(url string) error {
	pgUrl, err := pq.ParseURL(url)
	if err != nil {
		return err
	}

	var db *sql.DB
	db, err = sql.Open("postgres", pgUrl)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	r.db = db
	return nil
}
