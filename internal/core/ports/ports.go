package ports

import (
	"CTodo/internal/core/domain"
	"CTodo/pkg/utils"
)

type UserRepository interface {
	Register(username, email, password string) (*domain.User, error)
	Login(email, password string) (*utils.LoginResponse, error)
	ReadUser(id string) (*domain.User, error)
	UpdateUser(id, email, username, password string) error
	DeleteUser(id string) error
}
