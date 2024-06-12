package auth

import (
	"context"
	"database/sql"

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
		return utils.NewResponse(nil, err.Error(), 400)
	}

	result, err := repository.Login(ctx, db, param)

	if utils.FailOnError(err, "Failed to login") != nil {
		return utils.NewResponse(nil, "Failed to login", 400)
	}

	// #region generate jwt token
	token, err := generateJWTToken(result)
	if utils.FailOnError(err, "Failed to generate jwt token") != nil {
		return utils.NewResponse(nil, "Failed to generate jwt token", 400)
	}
	// #endregion

	return utils.NewResponse(token, "Login success", 200)
}
