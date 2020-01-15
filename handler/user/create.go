package user

import (
	. "RESTfulGo/handler"
	"RESTfulGo/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Create(g *gin.Context)  {
	var r  CreateRequest
	var err error

	if err := g.Bind(&r); err != nil {
		log.Debugf("参数效验失败 %s", r)
		//g.JSON(http.StatusOK, gin.H{"error": response.ErrBind})
		SendResponse(g, err, nil)
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
	if r.Username == "" {
		err = response.New(response.ErrUserNotFound,fmt.Errorf("参数不能为空"))
		log.Errorf(err, "参数不能为空")
		SendResponse(g, err, nil)
		return
	}

	if r.Password == "" {
		err = fmt.Errorf("password is empty")
		SendResponse(g, err, nil)
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
