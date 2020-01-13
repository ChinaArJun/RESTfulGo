package main

import (
	"RESTfulGo/router"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	// routers
	router.Load(g, middlewares...)

	err := http.ListenAndServe(":8080",g)
	log.Printf("开启服服务")
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
}
