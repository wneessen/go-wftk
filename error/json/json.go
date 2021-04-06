package json

import (
	wftkerr "github.com/wneessen/go-wftk/error"
)

type ErrorJson struct {
	Status     string
	StatusCode uint32
	Action     string
	ErrorMsg   string
}

// Return a standardized JSON error respone
func HttpJsonError(c uint32, a string, m string) *ErrorJson {
	errorJson := &ErrorJson{
		Status:     wftkerr.GetErrorString(c),
		StatusCode: c,
		Action:     a,
		ErrorMsg:   m,
	}

	return errorJson
}
