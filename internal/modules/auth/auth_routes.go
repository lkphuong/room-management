package auth

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/lkphuong/room-management/internal/middleware/auth"
)

func AuthRoutes(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", middleware.ValidateToken(), func(c *gin.Context) {
			Login(c)
		})
	}
}
