package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.SecureJsonPrefix(")]},\n")
	r.GET("/someJson", func(context *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// 将输出：while(1);["lena","austin","foo"]
		context.SecureJSON(http.StatusOK, names)
	})

	r.Run(":8080")
}
