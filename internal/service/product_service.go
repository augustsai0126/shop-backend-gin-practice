package service

import (
	"shop-backend-gin-practice/internal/domain"
	"shop-backend-gin-practice/internal/repository"
)

type ProductService interface {
	CreateProduct(product *domain.Product) error
	GetProduct(id uint) (*domain.Product, error)
	ListProducts(page, pageSize int) ([]*domain.Product, int64, error)
	UpdateProduct(product *domain.Product) error
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepo}
}

func (s *productService) CreateProduct(product *domain.Product) error {
	return s.productRepo.Create(product)
}

func (s *productService) GetProduct(id uint) (*domain.Product, error) {
	return s.productRepo.GetByID(id)
}

func (s *productService) ListProducts(page, pageSize int) ([]*domain.Product, int64, error) {
	return s.productRepo.List(page, pageSize)
}

func (s *productService) UpdateProduct(product *domain.Product) error {
	return s.productRepo.Update(product)
}
