package user

import (
	"RESTfulGo/handler"
	"RESTfulGo/model"
	"github.com/gin-gonic/gin"
	"log"
)

// @Summary
// @Description
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.GetUserResponse true "Get user"
// @Success 200 {object} handler.Result "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /v1/user [get]
func Get(g *gin.Context) {
	userId := g.Query("id")
	user, err := model.GetUser(userId)
	if err != nil {
		log.Println("GetUser Err:", err)
		return
	}
	handler.SendResponse(g, nil, user)
}
