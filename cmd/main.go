package main

import (
	"os"
	"shop-backend-gin-practice/config"
	"shop-backend-gin-practice/internal/handler"
	"shop-backend-gin-practice/internal/middleware"
	"shop-backend-gin-practice/internal/repository"
	"shop-backend-gin-practice/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// DB Connection
	dsn := config.GetPostgresDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	// GIN
	r := gin.Default()
	api := r.Group("/api")
	// Register
	api.POST("/register", userHandler.Register)
	// Login
	api.POST("/login", userHandler.Login)
	// GetMe
	api.GET("/me", middleware.JWTAuth(), userHandler.Me)

	// Admin API
	adminGroup := api.Group("/admin")
	adminGroup.Use(middleware.JWTAuth(), middleware.AdminAuth(userService))
	// New Category
	adminGroup.POST("/category/new", categoryHandler.New)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
