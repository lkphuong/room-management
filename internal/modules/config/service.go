package config

import (
	"context"
	"database/sql"

	"github.com/lkphuong/room-management/configs/http_code"
	"github.com/lkphuong/room-management/internal/utils"
)

var (
	repository Repository
)

type Service struct{}

func (s *Service) GetConfigStoreDetail(ctx context.Context, db *sql.DB, storeID string) *utils.Response {
	receipt, err := repository.ConfigStoreDetail(ctx, db, storeID)

	if utils.FailOnError(err, "Failed to get Store detail") != nil {
		return utils.NewResponse(nil, "Failed to get Store detail", http_code.BAD_REQUEST)
	}

	if receipt  == nil {
		return utils.NewResponse(nil, "Store detail not found", http_code.NOT_FOUND)
	}

	return utils.NewResponse(receipt, "Get Store detail success", http_code.OK)
}