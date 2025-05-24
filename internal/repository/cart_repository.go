package repository

import (
	"errors"
	"shop-backend-gin-practice/internal/domain"
	"time"

	"gorm.io/gorm"
)

type CartRepository interface {
	CreateCart(cart *domain.Cart) error
	GetCartByUserID(userID uint) (*domain.Cart, error)
	AddItem(userID, productID uint, quantity int) error
	RemoveItem(userID, productID uint) error
	UpdateItemQuantity(userID, productID uint, quantity int) error
	EmptyCart(cartID uint) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) CreateCart(cart *domain.Cart) error {
	return r.db.Create(cart).Error
}

func (r *cartRepository) GetCartByUserID(userID uint) (*domain.Cart, error) {
	var cart domain.Cart
	result := r.db.Preload("Items").Where("user_id = ?", userID).First(&cart)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil // 沒有找到回 nil
	}
	return &cart, result.Error
}

func (r *cartRepository) AddItem(userID, productID uint, quantity int) error {
	var cart domain.Cart
	result := r.db.First(&cart, "user_id =?", userID)
	if result.Error != nil {
		return result.Error
	}
	var item domain.CartItem
	result = r.db.First(&item, "cart_id =? AND product_id =?", cart.ID, productID)
	if result.Error != nil {
		item = domain.CartItem{
			CartID:    cart.ID,
			ProductID: productID,
			Quantity:  quantity,
		}
		return r.db.Create(&item).Error
	}
	item.Quantity += quantity
	err := r.db.Save(&item).Error
	if err != nil {
		return err
	}
	// 更新購物車的更新時間
	r.db.Model(&cart).Update("updated_at", time.Now()) // 更新 updated_at field
	return nil
}

func (r *cartRepository) RemoveItem(userID, productID uint) error {
	var cart domain.Cart
	result := r.db.First(&cart, "user_id =?", userID)
	if result.Error != nil {
		return result.Error
	}
	err := r.db.Where("cart_id =? AND product_id =?", cart.ID, productID).Delete(&domain.CartItem{}).Error
	if err != nil {
		return err
	}
	// 更新購物車的更新時間
	r.db.Model(&cart).Update("updated_at", time.Now()) // 更新 updated_at field
	return nil
}

func (r *cartRepository) UpdateItemQuantity(userID, productID uint, quantity int) error {
	var cart domain.Cart
	result := r.db.First(&cart, "user_id =?", userID)
	if result.Error != nil {
		return result.Error
	}
	var item domain.CartItem
	result = r.db.First(&item, "cart_id =? AND product_id =?", cart.ID, productID)
	if result.Error != nil {
		return result.Error
	}
	item.Quantity = quantity
	return r.db.Save(&item).Error
}

func (r *cartRepository) EmptyCart(cartID uint) error {
	err := r.db.Where("cart_id =?", cartID).Delete(&domain.CartItem{}).Error
	if err != nil {
		return err
	}
	// 更新購物車的更新時間
	r.db.Model(&domain.Cart{}).Where("id =?", cartID).Update("updated_at", time.Now()) // 更新 updated_at field
	return nil
}
