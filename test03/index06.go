package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	r := gin.Default()

	var msg struct {
		Name string `json:"user"`
		Message string
		Number int
	}

	// JSON
	r.GET("/someJson", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(context *gin.Context) {
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123

		context.JSON(http.StatusOK, msg)
	})

	// XML
	r.GET("/someXML", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	// Yaml
	r.GET("/someYAML", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	// protobuf
	r.GET("/someProtoBuf", func(context *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		data := &protoexample.Test{
			Label:            &label,
			Reps:             reps,
		}
		context.ProtoBuf(http.StatusOK, data)
	})

	r.Run(":8080")

}
