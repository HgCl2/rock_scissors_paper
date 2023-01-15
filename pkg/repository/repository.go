package repository

import "github.com/jmoiron/sqlx"

type Authorization interface{}

type Room interface{}

type Player interface{}

type Repository struct {
	Authorization
	Room
	Player
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
