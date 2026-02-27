package gozero

import (
	"encoding/json"
	"net/http"
)

// respondWithError 统一错误响应格式
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// respondWithError 写入错误响应
func respondWithError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	resp := ErrorResponse{
		Code:    code,
		Message: err.Error(),
	}

	_ = json.NewEncoder(w).Encode(resp)
}
