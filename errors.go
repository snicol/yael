package yael

import (
	"net/http"
)

const (
	// HTTP error codes
	BadRequest          = "bad_request"
	MethodNotAllowed    = "method_not_allowed"
	NotFound            = "not_found"
	UnprocessableEntity = "unprocessable_entity"
)

func StatusCode(e E) int {
	switch e.Code {
	case BadRequest:
		return http.StatusBadRequest
	case MethodNotAllowed:
		return http.StatusMethodNotAllowed
	case NotFound:
		return http.StatusNotFound
	case UnprocessableEntity:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusBadRequest
	}
}
