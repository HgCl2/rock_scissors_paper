package service

import "github.com/HgCl2/rock_scissors_paper/pkg/repository"

type Authorization interface{}

type Room interface{}

type Player interface{}

type Service struct {
	Authorization
	Room
	Player
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
