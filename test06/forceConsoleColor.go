package main

import "github.com/gin-gonic/gin"

func main() {
	// 强制控制台有颜色
	gin.ForceConsoleColor()

	route := gin.Default()

	route.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})

	route.Run(":8080")
}
