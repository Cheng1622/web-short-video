package app

import (
	"net/http"

	"github.com/Cheng1622/web-short-video/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code errcode.ResCode `json:"code"`
	Msg  interface{}     `json:"msg"`
	// omitempty 代表 data 为null 就不返回了
	Data interface{} `json:"data,omitempty"`
}

func ResponseError(c *gin.Context, code errcode.ResCode) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code errcode.ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: errcode.CodeSuccess,
		Msg:  errcode.CodeSuccess.Msg(),
		Data: data,
	})
}
