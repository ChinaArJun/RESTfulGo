package handler

import (
	"RESTfulGo/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int  `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

/*
* 返回统一的函数
*/
func SendResponse(g *gin.Context, err error, data interface{})  {
	// 解析错误
	code, message := response.DecodeErr(err)

	// 返回json
	g.JSON(http.StatusOK, Result{
		Code:    code,
		Message: message,
		Data:    data,
	})
}