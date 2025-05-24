package main

import (
	"os"
	"shop-backend-gin-practice/config"
	"shop-backend-gin-practice/internal/di"
	"shop-backend-gin-practice/internal/router"

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

	handlers := di.NewHandlerGroup(db)
	r := router.SetupRouter(handlers)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
