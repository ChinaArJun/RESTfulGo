package user

import (
	. "RESTfulGo/handler"
	"RESTfulGo/model"
	"RESTfulGo/pkg/auth"
	"RESTfulGo/pkg/response"
	"RESTfulGo/pkg/token"
	"github.com/gin-gonic/gin"
)

func Login(g *gin.Context)  {
	// binding the data with the user struct
	var r = model.UserModel{
		Username:g.PostForm("username"),
		Password:g.PostForm("password"),
	}

	// get mysql user
	user, err := model.GetUser(r.Username)
	if err != nil {
		SendResponse(g, response.ErrUserNotFound,nil)
		return
	}

	// 对比密码
	if err := auth.Compare(user.Password, r.Password); err != nil {
			SendResponse(g, response.ErrPasswordIncorrect, nil)
			return
	}
	
	// 登录并获取token get token jwt
	t, err := token.Sign(g, token.Context{Username:user.Username, ID:user.Id}, "")
	if err != nil {
		SendResponse(g, response.ErrToken, nil)
		return
	}
	// 返回生成的token信息
	SendResponse(g, nil,model.Token{Token: t})
}