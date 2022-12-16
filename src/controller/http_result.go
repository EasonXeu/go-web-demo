package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpResult struct {
	Code         string      `json:"Code"`
	ErrorMessage string      `json:"ErrorMessage"`
	Data         interface{} `json:"Data,omitempty"`
}

func SuccessWithoutData(c *gin.Context) {
	c.JSON(http.StatusOK, HttpResult{Code: "Success", ErrorMessage: "success", Data: nil})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, HttpResult{Code: "Success", ErrorMessage: "success", Data: data})
}

func Abort(c *gin.Context, errorCode ErrorCode, msg string) {
	c.JSON(errorCode.httpCode, HttpResult{Code: errorCode.errorCode, ErrorMessage: msg, Data: nil})
}

type ErrorCode struct {
	httpCode  int
	errorCode string
}

var (
	InvalidParam        ErrorCode = ErrorCode{400, "InvalidParam"}
	EntityNotExist      ErrorCode = ErrorCode{404, "EntityNotExist"}
	InternalServerError ErrorCode = ErrorCode{500, "InternalServerError"}
)
