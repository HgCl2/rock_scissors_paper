package repository

type Authorization interface{}

type Room interface{}

type Player interface{}

type Repository struct {
	Authorization
	Room
	Player
}

func NewRepository() *Repository {
	return &Repository{}
}
