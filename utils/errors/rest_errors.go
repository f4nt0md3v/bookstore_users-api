package errors

import "net/http"

type RestError struct {
	Error        string   `json:"error"`
	Message      string   `json:"message"`
	StatusCode   int      `json:"status_code"`
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Error:      "not_found",
		Message:    message,
		StatusCode: http.StatusNotFound,
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Error:      "internal_server_error",
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Error:      "bad_request",
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}
