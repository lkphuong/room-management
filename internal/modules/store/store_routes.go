package store

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/lkphuong/room-management/internal/middleware/auth"
)

func StoreRoutes(r *gin.Engine) {
	storeGroup := r.Group("/store")
	{
		storeGroup.GET("/", middleware.ValidateToken(), func(c *gin.Context) {
			GetAll(c)
		})
	}
}
