package services

import (
	"CTodo/internal/core/domain"
	"CTodo/internal/core/ports"
	"CTodo/pkg/utils"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) Register(username, email, password string) (*domain.User, error) {
	return u.repo.Register(username, email, password)
}

func (u *UserService) ReadUser(id string) (*domain.User, error) {
	return u.repo.ReadUser(id)
}

func (u *UserService) UpdateUser(id, email, username, password string) error {
	return u.repo.UpdateUser(id, email, username, password)
}

func (u *UserService) DeleteUser(id string) error {
	return u.repo.DeleteUser(id)
}

func (u *UserService) Login(email, password string) (*utils.LoginResponse, error) {
	return u.repo.Login(email, password)
}
