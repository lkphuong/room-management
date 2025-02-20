package store

import (
	"context"
	"database/sql"
	"fmt"

	receipt "github.com/lkphuong/room-management/internal/modules/receipt"
	room "github.com/lkphuong/room-management/internal/modules/room"
	"github.com/lkphuong/room-management/internal/utils"
)

var (
	repository        Repository
	roomRepository    room.Repository
	receiptRepository receipt.Repository
)

type Service struct{}

func (s *Service) GetStores(ctx context.Context, db *sql.DB) *utils.Response {

	stores, err := repository.GetStores(ctx, db)
	if utils.FailOnError(err, "Failed to get stores") != nil {
		return utils.NewResponse(nil, "Failed to get stores", 400)
	}

	rooms, err := roomRepository.GetRooms(ctx, db)
	if utils.FailOnError(err, "Failed to get rooms") != nil {
		return utils.NewResponse(nil, "Failed to get rooms", 400)
	}

	revenue, err := receiptRepository.RevenueStore(ctx, db)
	if utils.FailOnError(err, "Failed to get revenue") != nil {
		return utils.NewResponse(nil, "Failed to get revenue", 400)
	}

	var response []AllStoreResponse

	for _, store := range stores {
		roomActive := int64(0)
		for _, room := range rooms {
			fmt.Println(utils.ConvertTime(room.Start))
			if store.StoreID == room.StoreCode && utils.ConvertTime(room.Start) != "" {
				roomActive++
			}
		}
		var storeResponse = AllStoreResponse{
			StoreID:   store.StoreID,
			StoreName: store.StoreName,
			RoomCount: store.RoomCount,
			Active:    roomActive,
			InActive:  store.RoomCount - roomActive,
		}

		if len(revenue) > 0 {
			for _, rev := range revenue {
				if store.StoreID == rev.StoreCode {
					storeResponse.Revenue = rev.Revenue
					storeResponse.RevenueTmp = rev.RevenueTmp
				}
			}
		}

		response = append(response, storeResponse)
	}

	return utils.NewResponse(response, "")
}
