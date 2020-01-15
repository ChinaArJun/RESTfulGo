package user

import (
	. "RESTfulGo/handler"
	"RESTfulGo/model"
	"RESTfulGo/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"

)

/**
创建用户逻辑：
从 HTTP 消息体获取参数（用户名和密码）
参数校验
加密密码
在数据库中添加数据记录
返回结果（这里是用户名）
 */

func Create(g *gin.Context)  {
	var r  CreateRequest
	var err error

	if err := g.ShouldBindJSON(&r); err != nil {
		log.Debugf("参数效验失败 %s", r)
		//g.JSON(http.StatusOK, gin.H{"error": response.ErrBind})
		SendResponse(g, err, nil)
		return
	}
	
	// 效验参数
	if err := r.checkParam(); err != nil {
		SendResponse(g, err, nil)
		return
	}
	
	user := model.UserModel{
		Username:  r.Username,
		Password:  r.Password,
	}

	if err := user.Validate(); err != nil {
		SendResponse(g, response.ErrValidation, nil)
		return
	}
	
	// post获取到的参数
	username := g.Param("username")
	log.Infof("username = %s", username)

	// url获取到的参数
	desc := g.Query("desc")
	log.Infof("desc = %s", desc)

	contentType := g.GetHeader("Content-Type")
	log.Infof("Header contentType = %s", contentType)

	log.Debugf("username is [%s], password is [%s]", r.Username, r.Password)


	// 密码加密
	if err := user.Encrypt(); err != nil {
		log.Errorf(err, "Encrypt error")
		SendResponse(g, response.ErrEncrypt, nil)
		return
	}
	
	// 新增用户
	if err := user.Create(); err != nil {
		log.Errorf( err, "ErrDatabase")
		SendResponse(g, response.ErrDatabase, nil)
		return
	}

	if response.IsErrUserNotFound(err) {
		log.Debug("err type Is ErrUserNotFound")
	}

	rsp := CreateResponse{Username:username}

	//code, message := response.DecodeErr(err)
	//g.JSON(http.StatusOK, gin.H{"code":code, "message": message})
	SendResponse(g, nil, rsp)
}

func (r *CreateRequest) checkParam() error  {
	if r.Username == "" {
		return response.New(response.ErrUserNotFound,fmt.Errorf("参数不能为空"))
	}

	if r.Password == "" {
		return fmt.Errorf("password is empty")
	}
	return nil
}