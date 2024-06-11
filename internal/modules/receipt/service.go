package receipt

import (
	"context"
	"database/sql"

	"github.com/lkphuong/room-management/internal/utils"
)

var (
	repository Repository
)

type Service struct{}

func (s *Service) GetBillDetail(ctx context.Context, db *sql.DB, param ReceiptDetailParam) *utils.Response {

	if err := param.Validate(); err != nil {
		return utils.NewResponse(nil, err.Error(), 400)
	}

	receipt, err := repository.BillDetail(ctx, db, param.Store, param.Room)

	utils.FailOnError(err, "Failed to get bill detail")

	if len(receipt) == 0 {
		return utils.NewResponse(nil, "Bill detail not found", 404)
	}

	return utils.NewResponse(receipt, "Get bill detail success", 200)
}
