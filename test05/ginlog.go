package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	// 如何记录日志
	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})

	router.Run(":8080")
}