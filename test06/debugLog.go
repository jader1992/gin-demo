package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// 自定义日志路由

	route := gin.Default()

	// 设置日志路由格式
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	route.POST("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "foo")
	})

	route.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bar")
	})

	route.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	route.Run(":8080")
}
