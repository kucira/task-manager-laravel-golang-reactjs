package models

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`  // Data returned on success, omitted if empty
	Error   interface{} `json:"error,omitempty"` // Specific error info on failure, omitted if empty
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, APIResponse{
		Code:    code,
		Message: "success",
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string, err error) {
	c.JSON(code, APIResponse{
		Code:    code,
		Message: message,
		Error:   err.Error(),
	})
}
