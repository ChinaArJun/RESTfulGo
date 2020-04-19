package router

import (
	_ "RESTfulGo/docs" // 注册swag目录
	"RESTfulGo/handler/block"
	"RESTfulGo/handler/sd"
	"RESTfulGo/handler/user"
	"RESTfulGo/router/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
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

	// 设置请求ID生成全局中间件
	g.Use(middleware.RequestId())
	// 请求信息中间件  - 会消耗一些性能
	g.Use(middleware.Logging())

	g.POST("/login", user.Login)

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)       //创建用户
		u.DELETE("/:id", user.Delete) //删除用户
		u.PUT("/:id", user.Update)    // 更新用户
		u.GET("", user.List)          //用户列表
		u.GET("/:id", user.Get)       //获取指定用户的数据
	}

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	chain := g.Group("/chaincode")
	{
		chain.POST("/contract", block.ContractBlock)
	}

	return g
}
