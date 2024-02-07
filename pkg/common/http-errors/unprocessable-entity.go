package httpErrors

import "net/http"

func NewUnprocessableEntityError(message string) *HTTPError  {
	return NewHTTPError(http.StatusUnprocessableEntity, message)
}
