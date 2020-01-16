package user

import (
	"RESTfulGo/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Update(g *gin.Context) {

	g.JSON(http.StatusOK, handler.Result{
		Code:    200,
		Message: "update",
		Data:    nil,
	})
}
