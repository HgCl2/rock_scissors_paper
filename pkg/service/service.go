package service

import (
	app "github.com/HgCl2/rock_scissors_paper"
	"github.com/HgCl2/rock_scissors_paper/pkg/repository"
)

type Authorization interface {
	CreateUser(user app.Player) (int, error)
}

type Room interface{}

type Player interface{}

type Service struct {
	Authorization
	Room
	Player
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
