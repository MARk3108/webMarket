package services

import (
	"WebMarket/models"
	"WebMarket/repositories"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	Register(username, password string, isAdmin bool) error
	Login(username, password string) (string, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

var jwtSecret = []byte("marchekqwerty")

func (s *authService) Register(username, password string, isAdmin bool) error {
	user := &models.User{Username: username}
	if err := user.SetPassword(password); err != nil {
		return err
	}
	if isAdmin {
		user.Role = "admin"
	}

	return s.userRepo.Create(user)
}

func (s *authService) Login(username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil || !user.CheckPassword(password) {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  strconv.FormatUint(uint64(user.ID), 10),
		"username": username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString(jwtSecret)
}
