package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/someJson", func(context *gin.Context) {
		data := map[string]interface{} {
			"lang" : "Go语言",
			"Tag": "<br>",
		}

		context.AsciiJSON(http.StatusOK, data)
	})

	r.Run(":8080")
}
