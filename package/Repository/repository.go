package repository

import "github.com/jmoiron/sqlx"

type Submission interface {
}
type DataBase interface {
}

type Repository struct {
	Submission
	DataBase
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}