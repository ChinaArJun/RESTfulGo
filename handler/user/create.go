package user

import (
	"RESTfulGo/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"net/http"
)

func Create(g *gin.Context)  {

	var r struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var err error
	if err := g.Bind(&r); err != nil {
		log.Debugf("参数效验失败 %s", r)
		g.JSON(http.StatusOK, gin.H{"error": response.ErrBind})
		return
	}
	if r.Username == "" {
		err = response.New(response.ErrUserNotFound,fmt.Errorf("参数不能为空"))
		log.Errorf(err, "参数不能为空")
	}

	if response.IsErrUserNotFound(err) {
		log.Debug("err type Is ErrUserNotFound")
	}

	if r.Password == "" {
		err = fmt.Errorf("password is empty")
	}

	code, message := response.DecodeErr(err)
	g.JSON(http.StatusOK, gin.H{"code":code, "message": message})

}
