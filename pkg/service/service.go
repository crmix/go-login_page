package service

import (
	loginapi "login_api"
	"login_api/pkg/repository"
)

type Authorization interface {
	CreateUser(user loginapi.SignUp)(int, error)
	GenerateToken(email, password string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}