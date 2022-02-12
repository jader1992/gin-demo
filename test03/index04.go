package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.POST("/post", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page", "0")
		name := context.PostForm("name")
		message := context.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	route.Run(":8080")
}
