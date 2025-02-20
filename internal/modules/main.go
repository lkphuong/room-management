package modules

import (
	"database/sql"

	"github.com/go-playground/form"
	"github.com/go-playground/validator/v10"
	"github.com/volatiletech/sqlboiler/boil"

	sqlserver "github.com/lkphuong/room-management/configs/database"
)

var (
	db       *sql.DB
	validate *validator.Validate
	decoder  *form.Decoder
)

func GetDB() *sql.DB {
	db := sqlserver.ConnectionSqlServer()

	boil.SetDB(db)

	return db
}

func GetValidator() *validator.Validate {
	return validate
}
