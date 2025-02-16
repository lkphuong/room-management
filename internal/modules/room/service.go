package room

import (
	"context"
	"database/sql"

	"github.com/lkphuong/room-management/configs/http_code"
	receipt "github.com/lkphuong/room-management/internal/modules/receipt"
	"github.com/lkphuong/room-management/internal/utils"
	"github.com/lkphuong/room-management/internal/validations"
)

var (
	repository        Repository
	receiptRepository receipt.Repository
)

type Service struct{}

func (s *Service) GetRoomByStore(ctx context.Context, db *sql.DB, store string, user utils.JwtPayload) *utils.Response {
	errMsg := validations.ValidateUserInStore(store, user)
	if errMsg != nil {
		return utils.NewResponse(nil, *errMsg, http_code.BAD_REQUEST)
	}

	var rooms []RoomResponse

	rooms, err := repository.GetRoomsByStore(ctx, db, store)
	if utils.FailOnError(err, "Failed to get rooms") != nil {
		return utils.NewResponse(nil, "Failed to get rooms", http_code.BAD_REQUEST)
	}

	revenue, err := receiptRepository.RevenueRoom(ctx, db, store)
	if utils.FailOnError(err, "Failed to get revenue") != nil {
		return utils.NewResponse(nil, "Failed to get revenue", http_code.BAD_REQUEST)
	}

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
		if utils.FailOnError(err, "Failed to convert start date") != nil {
			return utils.NewResponse(nil, "Failed to convert start date", http_code.BAD_REQUEST)
		}

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
