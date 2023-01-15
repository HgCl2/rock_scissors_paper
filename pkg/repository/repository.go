package repository

import (
	app "github.com/HgCl2/rock_scissors_paper"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user app.Player) (int, error)
}

type Room interface{}

type Player interface{}

type Repository struct {
	Authorization
	Room
	Player
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
