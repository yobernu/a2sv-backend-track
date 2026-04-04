package services

import (
	"fmt"
	"go-auth/models"
	"go-auth/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	store *models.UserStore
}

func NewUserService(store *models.UserStore) *UserService {
	return &UserService{store: store}
}

func (s *UserService) Register(email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		12,
	)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	if err := s.store.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(email, password string) (*utils.TokenPair, error) {
	user, err := s.store.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	tokens, err := utils.GenerateTokenPair(user)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
