package service

import (
	app "github.com/HgCl2/rock_scissors_paper"
	"github.com/HgCl2/rock_scissors_paper/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user app.Player) (int, error) {
	return s.repo.CreateUser(user)
}
