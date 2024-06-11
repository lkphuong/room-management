package receipt

import (
	"database/sql"

	"github.com/go-playground/form"
	"github.com/go-playground/validator/v10"
	"github.com/lkphuong/room-management/internal/modules"
)

var (
	validation *validator.Validate
	_decoder   *form.Decoder
	db         *sql.DB
)

func init() {
	db = modules.GetDB()
	validation = validator.New()
}
