package server

import (
	"fmt"
	"net/http"
)

type HttpErrorResponse struct {
	StatusText string `json:"status_text"`
	StatusCode int    `json:"status_code"`
}

func NewError(statusCode int) *HttpErrorResponse {
	return &HttpErrorResponse{
		StatusText: http.StatusText(statusCode),
		StatusCode: statusCode,
	}
}

func (e *HttpErrorResponse) String() string {
	return fmt.Sprintf("%s %d", e.StatusText, e.StatusCode)
}
