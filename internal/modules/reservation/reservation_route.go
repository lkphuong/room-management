package reservation

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/lkphuong/room-management/internal/middleware/auth"
)

func ReservationRoutes(r *gin.RouterGroup) {
	reservationGroup := r.Group("/reservations")
	{
		reservationGroup.GET("/", middleware.ValidateToken(), func(c *gin.Context) {
			GetAll(c)
		})
	}
}