package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	SuccessCode uint = 10_000
	SuccessMsg       = "success"

	DefaultFailCode uint = 10_500
	DefaultFailMsg       = "fail"

	NotFoundCode uint = 10_404
	NotFoundMsg       = "not found"

	AuthFailCode uint = 10_401
	AuthFailMsg       = "auth failed"

	ErrorEntityCode uint = 10_422
	ErrorEntityMsg       = "unprocessable entity"

	AddFailCode    uint = 10_600
	AddFailMsg          = "add failed"
	SearchFailCode uint = 10_700
	SearchFailMsg       = "search failed"
	UpdateFailCode uint = 10_800
	UpdateFailMsg       = "update failed"
	DeleteFailCode uint = 10_900
	DeleteFailMsg       = "delete failed"
)

type Response struct {
	Code uint        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Result(code uint, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Ok(c *gin.Context) {
	Result(SuccessCode, SuccessMsg, nil, c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SuccessCode, message, nil, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SuccessCode, SuccessMsg, data, c)
}

func OkWithDetailed(msg string, data interface{}, c *gin.Context) {
	Result(SuccessCode, msg, data, c)
}

func Fail(c *gin.Context) {
	Result(DefaultFailCode, DefaultFailMsg, nil, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(DefaultFailCode, msg, nil, c)
}

func FailWithDetailed(code uint, msg string, c *gin.Context) {
	Result(code, msg, nil, c)
}
