package middleware

import (
	"RESTfulGo/handler"
	"RESTfulGo/pkg/response"
	"RESTfulGo/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(g *gin.Context) {
		if _, err := token.ParseRequest(g); err != nil {
			// 没有token，拦截请求
			handler.SendResponse(g, response.ErrTokenInvalid, nil)
			g.Abort()
			return
		}
		 g.Next()
	}
}