package receipt

import (
	"github.com/gin-gonic/gin"
	"github.com/lkphuong/room-management/internal/utils"
)

var (
	service *Service
)

func ReceiptDetail(c *gin.Context) {
	r := c.Request

	ctx := r.Context()

	var p ReceiptDetailParam

	room := c.Query("room")
	store := c.Query("store")

	p.Room = room
	p.Store = store

	result := service.GetBillDetail(ctx, db, p)

	utils.JSONResponse(*result, c)
}
