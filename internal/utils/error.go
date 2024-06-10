package utils

import "log"

func FailOnError(err error, msg string) *Response {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		return NewResponse(nil, msg, 400)
	}
	return nil
}
