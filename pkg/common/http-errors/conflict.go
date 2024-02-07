package httpErrors

import "net/http"

func NewConflictError(message string) *HTTPError {
	return NewHTTPError(http.StatusConflict, message)
}
