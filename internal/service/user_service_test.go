package service

import (
	"errors"
	"shop-backend-gin-practice/internal/domain"
	"testing"
)

type MockUserRepository struct {
	users map[string]*domain.User
}

func (m *MockUserRepository) Create(user *domain.User) error {
	if _, ok := m.users[user.Username]; ok {
		return errors.New("user already exists")
	}
	m.users[user.Username] = user
	return nil
}

func (m *MockUserRepository) GetByUsername(username string) (*domain.User, error) {
	user, ok := m.users[username]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// 實際測試
func TestUserService_Register_Success(t *testing.T) {
	mockRepo := &MockUserRepository{
		users: make(map[string]*domain.User),
	}
	service := NewUserService(mockRepo)
	err := service.Register("testuser", "test@example.com", "123456")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if mockRepo.users["testuser"] == nil {
		t.Errorf("Expected user to be created, but got nil")
	}
}

func TestUserService_Register_UserAlreadyExists(t *testing.T) {
	mockRepo := &MockUserRepository{
		users: make(map[string]*domain.User),
	}
	mockRepo.users["testuser"] = &domain.User{Username: "testuser"}
	service := NewUserService(mockRepo)
	err := service.Register("testuser", "test2@example.com", "abcdef")
	if err == nil {
		t.Errorf("expected error for duplicate username, got nil")
	}
	if err.Error() != "username already exists" {
		t.Errorf("unexpected error: %v", err)
	}
}
