package services

import (
	"CTodo/internal/core/domain"
	"CTodo/internal/core/ports"
	"CTodo/pkg/utils"
)

// Service contains repository as ports to DB
type UserService struct {
	repo ports.UserRepository
}

// New NewUserService instance
func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// Service Registration
func (u *UserService) Register(username, email, password string) (*domain.User, error) {
	return u.repo.Register(username, email, password)
}

// Service Login
func (u *UserService) Login(email, password string) (*utils.LoginResponse, error) {
	return u.repo.Login(email, password)
}

// Getting service user instance
func (u *UserService) GetUser(id string) (*domain.User, error) {
	return u.repo.GetUser(id)
}

// Updating service user instance
func (u *UserService) UpdateUser(id, email, username, password string) error {
	return u.repo.UpdateUser(id, email, username, password)
}

// Deleting service user instance
func (u *UserService) DeleteUser(id string) error {
	return u.repo.DeleteUser(id)
}
