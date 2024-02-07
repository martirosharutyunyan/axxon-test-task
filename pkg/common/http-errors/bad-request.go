package httpErrors

import "net/http"

func NewBadRequestError(message string) *HTTPError {
	return NewHTTPError(http.StatusBadRequest, message)
}
