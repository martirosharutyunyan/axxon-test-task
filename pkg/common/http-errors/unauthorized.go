package httpErrors

import "net/http"

func NewUnauthorizedError(message string) *HTTPError {
	return NewHTTPError(http.StatusUnauthorized, message)
}
