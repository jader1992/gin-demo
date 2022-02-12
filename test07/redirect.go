package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()

	route.GET("/test", func(context *gin.Context) {
		context.Request.URL.Path = "/test2"
		route.HandleContext(context)
	})

	route.GET("/test2", func(context *gin.Context) {
		context.JSON(200, gin.H{"hello": "world"})
	})

	route.GET("/test1", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	route.Run(":8080")
}
