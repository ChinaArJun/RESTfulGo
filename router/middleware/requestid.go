package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

// 全局中间件 返回全局请求ID
func RequestId() gin.HandlerFunc {
	return func(g *gin.Context) {
		requestId := g.Request.Header.Get("X-Request-Id")

		if requestId == "" {
			u4 := uuid.NewV4()
			requestId = u4.String()
		}

		g.Set("X-Request-Id", requestId)

		g.Writer.Header().Set("X-Request-Id", requestId)
		g.Next()
	}
}
