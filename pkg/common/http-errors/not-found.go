package httpErrors

import "net/http"

func NewNotFoundError(message string) *HTTPError {
	return NewHTTPError(http.StatusNotFound, message)
}
