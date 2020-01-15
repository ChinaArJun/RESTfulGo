package user

import (
	"RESTfulGo/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(g *gin.Context)  {

	g.JSON(http.StatusOK, handler.Result{
		Code:    200,
		Message: "list",
		Data:    nil,
	})
}