package service

import (
	"shop-backend-gin-practice/internal/domain"
	"shop-backend-gin-practice/internal/repository"
)

type CategoryService interface {
	New(name, description string) error
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo}
}

func (s *categoryService) New(name, description string) error {
	category := &domain.Category{
		Name:        name,
		Description: description,
	}
	return s.categoryRepo.Create(category)
}
