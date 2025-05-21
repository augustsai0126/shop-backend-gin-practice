package middleware

import (
	"net/http"
	"shop-backend-gin-practice/internal/service"

	"github.com/gin-gonic/gin"
)

func AdminAuth(userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 檢查是否為管理員
		userIDValue, exists := c.Get("user_id")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
			return
		}
		userID := userIDValue.(uint)
		isAdmin, err := userService.IsAdmin(userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "server error"})
			return
		}
		if !isAdmin {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "not admin"})
			return
		}
		c.Next()
	}
}
