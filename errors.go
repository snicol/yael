// Package yael provides a structured error type that conforms to the errors.Wrapper
// and error interfaces.
package yael

import (
	"net/http"
)

const (
	// BadRequest is the error code which equates to a HTTP 400 Bad Request.
	BadRequest = "bad_request"
	// Unauthorized is the error code which equates to a 401 Unauthorized.
	Unauthorized = "unauthorized"
	// Forbidden is the error code which equates to a 403 Forbidden.
	Forbidden = "forbidden"
	// NotFound is the error code which equates to a 404 Not Found.
	NotFound = "not_found"
	// MethodNotAllowed is the error code which equates to a 405 Method Not Allowed.
	MethodNotAllowed = "method_not_allowed"
	// Conflict is the error code which equates to a 409 Conflict.
	Conflict = "conflict"
	// NotImplemented is the error code which equates to a 501 Not Implemented.
	NotImplemented = "not_implemented"
	// ServiceUnavailable is the error code which equates to a 503 Service Unavailable.
	ServiceUnavailable = "service_unavailable"
)

//nolint:gochecknoglobals // lookup table avoids a high-complexity switch in StatusCode
var statusCodes = map[string]int{
	BadRequest:         http.StatusBadRequest,
	Unauthorized:       http.StatusUnauthorized,
	Forbidden:          http.StatusForbidden,
	NotFound:           http.StatusNotFound,
	MethodNotAllowed:   http.StatusMethodNotAllowed,
	Conflict:           http.StatusConflict,
	NotImplemented:     http.StatusNotImplemented,
	ServiceUnavailable: http.StatusServiceUnavailable,
}

// StatusCode returns the HTTP status code for the given error.
func StatusCode(e E) int {
	if code, ok := statusCodes[e.Code]; ok {
		return code
	}

	return http.StatusBadRequest
}
