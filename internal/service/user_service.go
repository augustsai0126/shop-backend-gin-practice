package service

import (
	"errors"
	"shop-backend-gin-practice/internal/domain"
	"shop-backend-gin-practice/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(username, email, password string) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) Register(username, email, password string) error {
	// 檢查是否有重複帳號
	u, _ := s.userRepo.GetByUsername(username)
	if u != nil {
		return errors.New("username already exists")
	}
	// 密碼加密
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// 創建用戶
	user := &domain.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hash),
	}
	return s.userRepo.Create(user)
}
