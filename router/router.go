package router

import (
	"RESTfulGo/handler/sd"
	"RESTfulGo/router/middleware"
	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine  {
	// 设置请求的Header 参数
	// 在处理某些请求时可能因为程序 bug 或者其他异常情况导致程序 panic，
	// 这时候为了不影响下一次请求的调用，需要通过 gin.Recovery()来恢复 API 服务器
	g.Use(gin.Recovery())
	// 强制浏览器不使用缓存
	g.Use(middleware.NoCache)
	// 浏览器跨域 OPTIONS 请求设置
	g.Use(middleware.Options)
	// 一些安全设置
	g.Use(middleware.Secure)
	g.Use(mw...)

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}


	return g
}
