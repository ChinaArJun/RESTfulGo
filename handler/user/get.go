package user

import (
	"RESTfulGo/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(g *gin.Context)  {


	g.JSON(http.StatusOK, handler.Result{
		Code:    200,
		Message: "get",
		Data:    nil,
	})
}