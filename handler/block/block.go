package block

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ContractBlock(g *gin.Context)  {
	// 返回json
	g.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": "testok",
		"msg": "ok",
	})
	//handler.SendResponse(g, nil, gin.H{
	//	"code": 1,
	//	"data": "testok",
	//	"msg": "ok",
	//})
}