package httpErrors

import "net/http"

type HTTPError struct {
	Message    string
	StatusCode int
	Status     string
}

func (httpError *HTTPError) Error() string {
	return httpError.Message
}

func NewHTTPError(statusCode int, errorMessage string) *HTTPError {
	return &HTTPError{StatusCode: statusCode, Message: errorMessage, Status: http.StatusText(statusCode)}
}
