package main

import (
	"RESTfulGo/config"
	"RESTfulGo/router"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
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

	g := gin.New()


	middlewares := []gin.HandlerFunc{}

	// routers
	router.Load(g, middlewares...)


	// 异步协程运行API服务检查
	go func() {
		err := pingServer()
		if err != nil {
			log.Printf("pingServer error")
		}
		// 检查服务成功
		log.Print("The router been deployed successfully!")
	}()

	err := http.ListenAndServe(viper.GetString("addr"),g)
	log.Printf("开启服服务")
	if err != nil {
		log.Printf("error: %s", err.Error())
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
		// 延迟1秒
		log.Print("time sleep 1 ")
		time.Sleep(time.Second)
	}
	return errors.New("服务检查错误")
}
