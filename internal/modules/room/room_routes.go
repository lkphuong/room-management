package room

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/lkphuong/room-management/internal/middleware/auth"
)

func RoomRoutes(r *gin.Engine) {
	roomGroup := r.Group("/room")
	{
		roomGroup.GET("/", middleware.ValidateToken(), func(c *gin.Context) {
			RoomStatus(c)
		})
	}
}
