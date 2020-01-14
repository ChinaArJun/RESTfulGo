package main

import (
	"RESTfulGo/config"
	"RESTfulGo/router"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/lexkong/log"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config","c","", "apiserver config file path.")
)

func main() {

	// init config
	if err:= config.Init(*cfg); err != nil {
		panic(err)
	}

	// 配置文件设置gin运行模式
	gin.SetMode(viper.GetString("runmode"))

	// 测试日志打印转存效果
	//testLog()

	g := gin.New()


	middlewares := []gin.HandlerFunc{}

	// routers
	router.Load(g, middlewares...)


	// 异步协程运行API服务检查
	go func() {
		err := pingServer()
		if err != nil {
			log.Fatal("pingServer error", err)
		}
		// 检查服务成功
		log.Info("The router been deployed successfully!")
	}()

	log.Infof("Start to requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func testLog()  {
	for {
		// 延迟1秒
		log.Info("2333333333333333333333333333233333333333333333333333333323333333333333333333333333332333333333333333333333333333")
		time.Sleep(time.Second *2)
	}
}

// API 服务器健康状态自检
func pingServer() error {
	for i := 9; i < 10 ; i++  {
		// ping the server By Get request to "health"
		res, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && res.StatusCode == http.StatusOK {
			return nil
		}
		log.Info("time sleep 1 ")
		//time.Sleep(time.Second)
	}
	return errors.New("服务检查错误")
}
