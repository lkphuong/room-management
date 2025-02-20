package store

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/lkphuong/room-management/internal/middleware/auth"
)

func StoreRoutes(r *gin.RouterGroup) {
	storeGroup := r.Group("/stores")
	{
		storeGroup.GET("/", middleware.ValidateToken(), func(c *gin.Context) {
			GetAll(c)
		})

		storeGroup.GET("/my-stores", middleware.ValidateToken(), func(c *gin.Context) {
			GetMyStores(c)
		})
	}
}
