package user

import (
	"RESTfulGo/handler"
	"RESTfulGo/model"
	"RESTfulGo/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"strconv"
)

// @Summary Delete new user to the database
// @Description Delete a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Delete a new user"
// @Success 200 {object} handler.Result "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /v1/user [delete]
func Delete(g *gin.Context) {
	userId, _ := strconv.ParseUint(g.Param("id"), 10, 64)
	if err := model.DeleteUser(userId); err != nil {
		log.Fatal("delete user err:", err)
		handler.SendResponse(g, response.ErrDatabase, nil)
		return
	}
	handler.SendResponse(g, nil, nil)
}
