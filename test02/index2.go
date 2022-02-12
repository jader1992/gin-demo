package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	route.LoadHTMLGlob("templates/*/*")

	route.GET("/posts/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "posts",
		})
	})

	route.GET("/users/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "users",
		})
	})

	route.Run(":8080")
}
