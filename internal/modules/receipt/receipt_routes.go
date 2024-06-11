package receipt

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/lkphuong/room-management/internal/middleware/auth"
)

func ReceiptRoutes(r *gin.RouterGroup) {
	receiptGroup := r.Group("/receipts")
	{
		receiptGroup.GET("/", middleware.ValidateToken(), func(c *gin.Context) {
			ReceiptDetail(c)
		})
	}
}
