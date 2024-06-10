package store

import (
	"context"
	"database/sql"

	room "github.com/lkphuong/room-management/internal/modules/room"
	"github.com/lkphuong/room-management/internal/utils"
)

var (
	repository     Repository
	roomRepository room.Repository
)

func init() {
	repository = Repository{}
	roomRepository = room.Repository{}
}

type Service struct{}

func (s *Service) GetStores(ctx context.Context, db *sql.DB) *utils.Response {
	stores, err := repository.GetStores(ctx, db)
	utils.FailOnError(err, "Failed to get stores")

	rooms, err := roomRepository.GetRooms(ctx, db)
	utils.FailOnError(err, "Failed to get rooms")

	var response []AllStoreResponse

	for _, store := range stores {
		roomActive := int64(0)
		for _, room := range rooms {
			if store.StoreID == room.StoreCode && len(room.Start) > 0 {
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

		response = append(response, storeResponse)
	}

	return utils.NewResponse(response, "")
}
