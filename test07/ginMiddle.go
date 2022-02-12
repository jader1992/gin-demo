package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc  {
	return func(context *gin.Context) {
		t := time.Now()

		// 设置example变量
		context.Set("example", "12345")

		// 请求前
		context.Next()

		// 请求后
		latency := time.Since(t)
		log.Println(latency)

		// 获取发送的status
		status := context.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.Default()
	r.Use(Logger())

	r.GET("/test", func(context *gin.Context) {
		example := context.MustGet("example").(string)
		log.Println(example)
		context.JSON(200, gin.H{
			"example": example,
		})
	})

	r.Run(":8080")
}
