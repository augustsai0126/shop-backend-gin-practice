package main

import (
	"os"
	"shop-backend-gin-practice/config"
	"shop-backend-gin-practice/internal/handler"
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
	// GIN
	r := gin.Default()
	api := r.Group("/api")
	// Register
	api.POST("/register", userHandler.Register)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
