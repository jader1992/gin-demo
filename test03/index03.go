package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 提供 unicode 实体
	r.GET("/json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"html": "<b>hello world</b>",
		})
	})

	// 提供字面字符
	r.GET("/purejson", func(context *gin.Context) {
		context.PureJSON(http.StatusOK, gin.H{
			"html": "<b>hello world</b>",
		})
	})

	r.Run(":8080")
}
