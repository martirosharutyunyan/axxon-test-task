package httpErrors

import "net/http"

func NewForbiddenError(message string) *HTTPError {
	return NewHTTPError(http.StatusForbidden, message)
}
