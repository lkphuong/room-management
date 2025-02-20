package utils

import (
	"fmt"

	"github.com/lkphuong/room-management/configs/http_code"
)

func FailOnError(err error, msg string) *Response {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
		return NewResponse(nil, msg, http_code.BAD_REQUEST)
	}
	return nil
}
