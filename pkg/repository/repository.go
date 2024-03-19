package repository

import (
	loginapi "login_api"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user loginapi.SignUp) (int, error)
	GetUser(email, password string) (loginapi.SignUp, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}