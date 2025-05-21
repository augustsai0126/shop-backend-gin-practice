package service

import (
	"errors"
	"shop-backend-gin-practice/config"
	"shop-backend-gin-practice/internal/domain"
	"shop-backend-gin-practice/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(username, email, password string) error
	Login(username, password string) (string, error)
	GetUserByID(userID uint) (*domain.User, error)
	IsAdmin(userID uint) (bool, error)
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

func (s *userService) Login(username, password string) (string, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil || user == nil {
		return "", errors.New("invalid username or password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// 生成 JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"user_name": user.Username,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})
	jwtSecret := []byte(config.GetJWTSecret())
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *userService) GetUserByID(userID uint) (*domain.User, error) {
	return s.userRepo.GetByID(userID)
}

func (s *userService) IsAdmin(userID uint) (bool, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return false, err
	}
	return user.IsAdmin, nil
}
