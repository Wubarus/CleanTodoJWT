package ports

import (
	"github.com/Wubarus/CleanTodoJWT/internal/core/domain"
	"github.com/Wubarus/CleanTodoJWT/pkg/utils"
)

// Ports are guarantee that its instance has listed methods

// UserRepository designed to manage user interaction
type UserRepository interface {
	Register(username, email, password string) (*domain.User, error)
	Login(email, password string) (*utils.LoginResponse, error)
	GetUser(id string) (*domain.User, error)
	UpdateUser(id, email, username, password string) error
	DeleteUser(id string) error
}
