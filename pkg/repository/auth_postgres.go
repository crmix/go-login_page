package repository

import (
	"fmt"
	loginapi "login_api"
    "github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user loginapi.SignUp) (int, error) {
	var id int
	query :=fmt.Sprintf("INSERT INTO %s (name, email, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Email, user.Password)
	if err :=row.Scan(&id); err!=nil {
		return 0, err
	}
	return id, nil
}

func(r *AuthPostgres) GetUser(email, password string) (loginapi.SignUp, error){
	var user loginapi.SignUp
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 and password_hash=$2", usersTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}