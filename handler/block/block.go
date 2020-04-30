package block

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ContractBlock(g *gin.Context) {
	log.Println("params:", g.Params)
	log.Println("params:", g.PostForm("id"))
	// 返回json
	g.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": "testok",
		"msg":  "ok",
	})
	//handler.SendResponse(g, nil, gin.H{
	//	"code": 1,
	//	"data": "testok",
	//	"msg": "ok",
	//})
}

func PatchContractBlock(g *gin.Context) {
	log.Println("params:", g.Params)
	log.Println("params:", g.PostForm("id"))
	// 返回json
	g.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": "testok",
		"msg":  "ok",
	})
	//handler.SendResponse(g, nil, gin.H{
	//	"code": 1,
	//	"data": "testok",
	//	"msg": "ok",
	//})
}
