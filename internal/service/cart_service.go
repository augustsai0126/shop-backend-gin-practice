package service

import (
	"shop-backend-gin-practice/internal/domain"
	"shop-backend-gin-practice/internal/repository"
)

type CartService interface {
	AddToCart(userID uint, productID uint, quantity int) error
	GetCartByUserID(userID uint) (*domain.Cart, error)
	RemoveFromCart(userID uint, productID uint) error
	EmptyCart(userID uint) error
}

type cartService struct {
	cartRepo repository.CartRepository
}

func NewCartService(cartRepo repository.CartRepository) CartService {
	return &cartService{cartRepo}
}

func (s *cartService) AddToCart(userID uint, productID uint, quantity int) error {
	cart, err := s.cartRepo.GetCartByUserID(userID)
	if err != nil || cart == nil {
		cart = &domain.Cart{
			UserID: userID,
		}
		if err := s.cartRepo.CreateCart(cart); err != nil {
			return err
		}
	}
	return s.cartRepo.AddItem(userID, productID, quantity)
}

func (s *cartService) GetCartByUserID(userID uint) (*domain.Cart, error) {
	return s.cartRepo.GetCartByUserID(userID)
}

func (s *cartService) RemoveFromCart(userID uint, productID uint) error {
	return s.cartRepo.RemoveItem(userID, productID)
}

func (s *cartService) EmptyCart(userID uint) error {
	cart, err := s.cartRepo.GetCartByUserID(userID)
	if err != nil || cart == nil {
		return nil
	}
	return s.cartRepo.EmptyCart(cart.ID)
}
