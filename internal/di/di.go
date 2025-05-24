package di

import (
	"shop-backend-gin-practice/internal/handler"
	"shop-backend-gin-practice/internal/repository"
	"shop-backend-gin-practice/internal/service"

	"gorm.io/gorm"
)

func NewHandlerGroup(db *gorm.DB) *handler.HandlerGroup {
	// repository
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	cartRepo := repository.NewCartRepository(db)
	// service
	userService := service.NewUserService(userRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	productService := service.NewProductService(productRepo)
	cartService := service.NewCartService(cartRepo)
	// handler
	userHandler := handler.NewUserHandler(userService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	productHandler := handler.NewProductHandler(productService)
	cartHandler := handler.NewCartHandler(cartService)

	return &handler.HandlerGroup{
		User:     userHandler,
		Category: categoryHandler,
		Product:  productHandler,
		Cart:     cartHandler,
	}
}
