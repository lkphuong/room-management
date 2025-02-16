package receipt

import (
	"context"
	"database/sql"

	"github.com/lkphuong/room-management/configs/http_code"
	"github.com/lkphuong/room-management/internal/utils"
	"github.com/lkphuong/room-management/internal/validations"
)

var (
	repository Repository
)

type Service struct{}

func (s *Service) GetBillDetail(ctx context.Context, db *sql.DB, param ReceiptDetailParam, user utils.JwtPayload) *utils.Response {

	if err := param.Validate(); err != nil {
		return utils.NewResponse(nil, err.Error(), http_code.BAD_REQUEST)
	}

	errMsg := validations.ValidateUserInStore(param.Store, user)
	if errMsg != nil {
		return utils.NewResponse(nil, *errMsg, http_code.BAD_REQUEST)
	}

	receipt, err := repository.BillDetail(ctx, db, param.Store, param.Room)

	if utils.FailOnError(err, "Failed to get bill detail") != nil {
		return utils.NewResponse(nil, "Failed to get bill detail", http_code.BAD_REQUEST)
	}

	if len(receipt) == 0 {
		return utils.NewResponse(nil, "Bill detail not found", http_code.NOT_FOUND)
	}

	return utils.NewResponse(receipt, "Get bill detail success", http_code.OK)
}
