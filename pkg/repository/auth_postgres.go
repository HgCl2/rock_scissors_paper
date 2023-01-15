package repository

import (
	app "github.com/HgCl2/rock_scissors_paper"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user app.Player) (int, error) {
	return 0, nil
}
