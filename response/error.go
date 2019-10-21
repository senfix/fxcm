package response

import "fmt"

type ApiError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func (t *ApiError) Error() string {
	return fmt.Sprintf("%v: %v", t.Code, t.Message)
}
