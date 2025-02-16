package auth

import (
	"context"
	"database/sql"

	"github.com/lkphuong/room-management/configs/http_code"
	"github.com/lkphuong/room-management/internal/utils"
)

var (
	repository Repository
)

func init() {
	repository = Repository{}
}

type Service struct{}

func (s *Service) Login(ctx context.Context, db *sql.DB, param LoginParam) *utils.Response {
	if err := param.Validate(); err != nil {
		return utils.NewResponse(nil, err.Error(), http_code.BAD_REQUEST)
	}

	result, err := repository.Login(ctx, db, param)

	if utils.FailOnError(err, "Failed to login") != nil || len(result) == 0 {
		return utils.NewResponse(nil, "Failed to login", http_code.BAD_REQUEST)
	}

	storeIDs := []string{}
	for _, r := range result {
		storeIDs = append(storeIDs, r.StoreID)
	}

	payload := JwtPayload{
		ID:       result[0].ID,
		Code:     result[0].Code,
		Name:     result[0].Name,
		StoreIDs: storeIDs,
	}

	// #region generate jwt token
	token, err := generateJWTToken(payload)
	if utils.FailOnError(err, "Failed to generate jwt token") != nil {
		return utils.NewResponse(nil, "Failed to generate jwt token", http_code.BAD_REQUEST)
	}
	// #endregion

	return utils.NewResponse(token, "Login success", http_code.OK)
}
