package auth

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

type AuthRepository struct{}

func (au *AuthRepository) Login(ctx context.Context, db *sql.DB, param LoginParam) (LoginResponse, error) {
	var result LoginResponse

	err := queries.Raw(fmt.Sprintf(LOGIN, param.Username, param.Password)).Bind(ctx, db, &result)

	if err != nil {
		return result, err
	}

	return result, nil
}
