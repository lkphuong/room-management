package room

import (
	"github.com/gin-gonic/gin"
	"github.com/lkphuong/room-management/internal/utils"
)

var (
	service *Service
)

func RoomStatus(c *gin.Context) {
	r := c.Request
	ctx := r.Context()

	store := c.Query("store")

	result := service.GetRoomByStores(ctx, db, store)

	utils.JSONResponse(*result, c)
}
