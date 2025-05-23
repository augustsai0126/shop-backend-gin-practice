package repository

import (
	"shop-backend-gin-practice/internal/domain"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *domain.Product) error
	GetAll() ([]*domain.Product, error)
	GetByID(id uint) (*domain.Product, error)
	Update(product *domain.Product) error
	List(page, pageSize int) ([]*domain.Product, int64, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Create(product *domain.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetAll() ([]*domain.Product, error) {
	var products []*domain.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) GetByID(id uint) (*domain.Product, error) {
	var product domain.Product
	result := r.db.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *productRepository) Update(product *domain.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) List(page, pageSize int) ([]*domain.Product, int64, error) {
	var products []*domain.Product
	var total int64
	r.db.Model(&domain.Product{}).Count(&total)
	offset := (page - 1) * pageSize
	result := r.db.Offset(offset).Limit(pageSize).Find(&products)
	return products, total, result.Error
}
