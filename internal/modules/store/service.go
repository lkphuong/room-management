package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lkphuong/room-management/configs/database"
	"github.com/lkphuong/room-management/configs/http_code"
	config "github.com/lkphuong/room-management/internal/modules/config"
	receipt "github.com/lkphuong/room-management/internal/modules/receipt"
	room "github.com/lkphuong/room-management/internal/modules/room"
	"github.com/lkphuong/room-management/internal/utils"
)

var (
	repository        Repository
	roomRepository    room.Repository
	receiptRepository receipt.Repository
	configRepository config.Repository
)

type Service struct{}

func (s *Service) GetStores(ctx context.Context, db *sql.DB) *utils.Response {
	
	stores, err := repository.GetStores(ctx, db)
	if utils.FailOnError(err, "Failed to get stores") != nil {
		return utils.NewResponse(nil, "Failed to get stores", http_code.BAD_REQUEST)
	}

	rooms, err := roomRepository.GetRooms(ctx, db)
	if utils.FailOnError(err, "Failed to get rooms") != nil {
		return utils.NewResponse(nil, "Failed to get rooms", http_code.BAD_REQUEST)
	}

	revenue, err := receiptRepository.RevenueStore(ctx, db)
	if utils.FailOnError(err, "Failed to get revenue") != nil {
		return utils.NewResponse(nil, "Failed to get revenue", http_code.BAD_REQUEST)
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

func (s *Service) GetMyStores(ctx context.Context, db *sql.DB, storeIDs []string) *utils.Response {
	configData, err := configRepository.ConfigStoreDetail(ctx,db,storeIDs[0])

	if utils.FailOnError(err, "Failed to get Store detail") != nil {
		return utils.NewResponse(nil, "Failed to get Store detail", http_code.BAD_REQUEST)
	}

	newDB := database.DynamicConnectionSqlServer(configData.Host,configData.Username,configData.Password,configData.Port,configData.Database)
	defer newDB.Close()

	stores, err := repository.GetStoreByIDs(ctx, storeIDs, newDB)
	if utils.FailOnError(err, "Failed to get stores") != nil {
		return utils.NewResponse(nil, "Failed to get stores", http_code.BAD_REQUEST)
	}

	rooms, err := roomRepository.GetRooms(ctx, newDB)
	if utils.FailOnError(err, "Failed to get rooms") != nil {
		return utils.NewResponse(nil, "Failed to get rooms", http_code.BAD_REQUEST)
	}

	revenue, err := receiptRepository.RevenueStore(ctx, newDB)
	if utils.FailOnError(err, "Failed to get revenue") != nil {
		return utils.NewResponse(nil, "Failed to get revenue", http_code.BAD_REQUEST)
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
