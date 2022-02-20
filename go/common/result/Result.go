package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpResult struct {
	Code string      `json:"Code"`
	Msg  string      `json:"Msg"`
	Data interface{} `json:"Data,omitempty"`
}

func SuccessWithoutData(c *gin.Context) {
	c.JSON(http.StatusOK, HttpResult{Code: "Success", Msg: "success", Data: nil})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, HttpResult{Code: "Success", Msg: "success", Data: data})
}

func Error(c *gin.Context, errorCode ErrorCode, msg string) {
	c.JSON(errorCode.httpCode, HttpResult{Code: errorCode.errorCode, Msg: msg, Data: nil})
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
