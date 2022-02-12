package main

import "github.com/gin-gonic/gin"

type myForm struct {
	Colors []string `form:"colors[]"`
}

func main() {

	route := gin.Default()
	route.POST("/post", func(context *gin.Context) {
		var fakeForm myForm
		context.ShouldBind(&fakeForm)
		context.JSON(200, gin.H{
			"color": fakeForm.Colors,
		})
	})
}
