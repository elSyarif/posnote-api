package helper

import (
	"github.com/gin-gonic/gin"
)

type responseSuccess struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

type responseSuccessWithData struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type responseError struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func HTTPResponseSuccess(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, responseSuccess{
		Status:  "success",
		Message: message,
	})
}

func HTTPResponseSuccessWithData(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, responseSuccessWithData{
		Status: "success",
		Data:   data,
	})
}

func HTTPResponseError(ctx *gin.Context, code int, status string, message string, err interface{}) {
	ctx.JSON(code, responseError{
		Status:  status,
		Message: message,
		Errors:  err,
	})
}
