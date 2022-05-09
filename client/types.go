package client

import (
	"fmt"
)

// HTTPError 请求错误响应
type HTTPError struct {
	HTTPCode int    `json:"httpCode"`
	Code     string `json:"code"`
	Message  string `json:"message"`
	Status   string `json:"status"`
	Err      string `json:"error"`
}

func (e *HTTPError) Error() string {
	detail := e.Err
	if len(detail) == 0 {
		detail = e.Status
	}
	return fmt.Sprintf("HTTPCode=%d, %s(%s): %s", e.HTTPCode, e.Message, e.Code, detail)
}
