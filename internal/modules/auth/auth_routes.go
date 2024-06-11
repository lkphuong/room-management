package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", func(c *gin.Context) {
			Login(c)
		})
	}
}
