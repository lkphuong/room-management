package room

import (
	"context"
	"database/sql"

	receipt "github.com/lkphuong/room-management/internal/modules/receipt"
	"github.com/lkphuong/room-management/internal/utils"
)

var (
	repository        Repository
	receiptRepository receipt.Repository
)

type Service struct{}

func (s *Service) GetRoomByStores(ctx context.Context, db *sql.DB, store string) *utils.Response {

	var rooms []RoomResponse

	rooms, err := repository.GetRoomsByStore(ctx, db, store)
	utils.FailOnError(err, "Failed to get rooms")

	revenue, err := receiptRepository.RevenueRoom(ctx, db, store)
	utils.FailOnError(err, "Failed to get revenue")

	var counter = 0
	var roomsResponse []RoomResponse
	for _, room := range rooms {
		if utils.ConvertTime(room.Start) != "" {
			counter++
		}
		var roomResponse = RoomResponse{
			StoreCode: room.StoreCode,
			RoomCode:  room.RoomCode,
			StoreName: room.StoreName,
		}

		start, err := utils.FormatDateString(room.Start)
		utils.FailOnError(err, "Failed to convert start date")

		roomResponse.Start = utils.ConvertTime(start)

		opened := utils.CalculateTime(start)
		roomResponse.Opened = opened

		if len(revenue) > 0 {
			for _, rev := range revenue {
				if room.RoomCode == rev.RoomCode {
					roomResponse.Revenue = rev.Revenue
					roomResponse.RevenueTmp = rev.RevenueTmp
				}
			}
		}

		roomsResponse = append(roomsResponse, roomResponse)
	}

	response := StatusResponse{
		Total:    int64(len(rooms)),
		Active:   int64(counter),
		InActive: int64(len(rooms) - counter),
		Data:     roomsResponse,
	}

	return utils.NewResponse(response, "")
}
