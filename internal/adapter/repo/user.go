package repo

import (
	"CTodo/internal/core/domain"
	"CTodo/pkg/utils"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lpernett/godotenv"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func (s *Storage) Register(username, email, password string) (*domain.User, error) {
	var user domain.User
	query := s.db.First(&user, "email = ?", email)
	if query.RowsAffected != 0 {
		return nil, errors.New("user at that email exists")
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, errors.New("cannot hash password, try again")
	}

	user = domain.User{
		Id:       uuid.New().String(),
		Email:    email,
		Username: username,
		Password: string(hashedPwd),
	}
	query = s.db.Create(&user)
	if query.RowsAffected == 0 {
		return nil, fmt.Errorf("user not saved")
	}

	return &user, nil
}

func (s *Storage) Login(email, password string) (*utils.LoginResponse, error) {
	var user domain.User
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	query := s.db.First(&user, "email = ?", email)
	if query.RowsAffected != 0 {
		return nil, errors.New("user not found")
	}

	accessToken, err := utils.GenerateAccessToken(user.Id, os.Getenv("JWT_SECRET"))
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.Id, os.Getenv("JWT_SECRET"))
	if err != nil {
		return nil, err
	}

	return &utils.LoginResponse{
		Id:           user.Id,
		Email:        user.Email,
		Username:     user.Username,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Storage) ReadUser(id string) (*domain.User, error) {
	var user domain.User

	query := s.db.First(&user, "id = ? ", id)
	if query.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (s *Storage) UpdateUser(id, email, username, password string) error {
	var user domain.User
	query := s.db.First(&user, "id = ?", id)
	if query.RowsAffected == 0 {
		return errors.New("user not found")
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return fmt.Errorf("password not hashed: %v", err)
	}

	updates := map[string]any{
		"email":    email,
		"password": string(hashedPwd),
		"username": username,
	}

	query = s.db.Model(&user).Where("id = ?", id).Updates(updates)
	if query.RowsAffected == 0 {
		return errors.New("unable to update user :(")
	}

	return nil
}

func (s *Storage) DeleteUser(id string) error {
	user := &domain.User{}
	query := s.db.Where("id = ?", id).Delete(&user)
	if query.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
