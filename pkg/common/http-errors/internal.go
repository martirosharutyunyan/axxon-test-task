package httpErrors

import "net/http"

func NewInternalServerError(message string) *HTTPError {
	return NewHTTPError(http.StatusInternalServerError, message)
}
