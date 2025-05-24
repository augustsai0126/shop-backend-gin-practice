package router

import (
	"shop-backend-gin-practice/internal/handler"
	"shop-backend-gin-practice/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.HandlerGroup) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	// member group
	member := api.Group("/member")
	member.POST("/register", h.User.Register)
	member.POST("/login", h.User.Login)
	member.GET("/me", middleware.JWTAuth(), h.User.Me)

	// admin group
	admin := api.Group("/admin")
	admin.Use(middleware.JWTAuth(), middleware.AdminAuth(h.UserService))
	// New Category
	admin.POST("/category/new", h.Category.New)
	// New Product
	admin.POST("/product/new", h.Product.CreateProduct)
	// Update Product
	admin.PUT("/product/:id", h.Product.UpdateProduct)

	// product group
	product := api.Group("/product")
	product.GET("/list", h.Product.ListProducts)
	product.GET("/:id", h.Product.GetProduct)

	// cart group
	cart := api.Group("/cart")
	cart.Use(middleware.JWTAuth())
	cart.POST("/add", h.Cart.AddToCart)
	cart.GET("/list", h.Cart.GetCart)
	cart.POST("/remove", h.Cart.RemoveFromCart)
	cart.DELETE("/empty", h.Cart.EmptyCart)

	return r
}
