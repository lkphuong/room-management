package auth

import (
	"context"
	"database/sql"

	"github.com/lkphuong/room-management/internal/utils"
)

var (
	_repository AuthRepository
)

func init() {
	_repository = AuthRepository{}
}

type Service struct{}

func (s *Service) Login(ctx context.Context, db *sql.DB, param LoginParam) *utils.Response {
	if err := param.Validate(); err != nil {
		return utils.NewResponse(nil, err.Error(), 400)
	}

	result, err := _repository.Login(ctx, db, param)

	utils.FailOnError(err, "Failed to login")

	// #region generate jwt token
	token, err := generateJWTToken(result)
	utils.FailOnError(err, "Failed to generate jwt token")
	// #endregion

	return utils.NewResponse(token, "Login success", 200)
}
