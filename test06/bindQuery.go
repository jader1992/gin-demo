package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type localPerson struct {
	Name string `form:"name"`
	Address string `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	route := gin.Default()
	route.GET("/testing", startPage)
	route.Run(":8080")
}

func startPage(context *gin.Context) {
	var person localPerson

	if context.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Address)
	}

	context.String(200, "Success")
}
