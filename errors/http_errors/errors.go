package http_errors

import (
	"encoding/json"
	"fmt"
)

type HttpErrorResponse struct {
	StatusText string `json:"status_text"`
	StatusCode int    `json:"status_code"`
}

func NewError(statusCode int, statusText string) *HttpErrorResponse {
	return &HttpErrorResponse{
		StatusText: statusText,
		StatusCode: statusCode,
	}
}

func (e *HttpErrorResponse) String() string {
	return fmt.Sprintf("%s %d", e.StatusText, e.StatusCode)
}

func (h HttpErrorResponse) AsJson() []byte {
	data, _ := json.Marshal(h)

	return data
}
