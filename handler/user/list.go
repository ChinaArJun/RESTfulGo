package user

import (
	"RESTfulGo/handler"
	"RESTfulGo/model"
	"github.com/gin-gonic/gin"
	"log"
)

// @Summary 获取用户列表
// @Description Delete a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} user.ListResponse "{"code":0,"message":"OK","data":[{"username":"kong"},{"username":"kong"}]}"
// @Router /v1/user [get]
func List(g *gin.Context) {
	log.Println("List users")
	var listRequest ListRequest
	if err := g.ShouldBind(&listRequest); err != nil{
		log.Fatal("err ShouldBind listRequest")
		return
	}
	userList, count, err := model.GetUserList(listRequest.Pagesize, listRequest.PageNum)
	if err != nil {
		log.Fatal("getUserList err:", err)
		handler.SendResponse(g, err, nil)
		return
	}
	handler.SendResponse(g, nil, ListResponse{
		TotalCount: count,
		UserList: userList,
	} )
}
