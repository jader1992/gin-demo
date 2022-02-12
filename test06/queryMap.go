package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.POST("/testing", func(context *gin.Context) {
		ids := context.QueryMap("ids")
		names := context.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})

	route.Run(":8080")
}