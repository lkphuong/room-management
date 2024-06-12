package utils

import (
	"fmt"
)

func FailOnError(err error, msg string) *Response {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
		return NewResponse(nil, msg, 400)
	}
	return nil
}
