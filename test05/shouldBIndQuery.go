package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// ShouldBindQuery 函数只绑定 url 查询参数而忽略 post 数据
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":8080")
}

type Person struct {
	Name string `form:"name"`
	Address string `form:"address"`
}

func startPage(c *gin.Context)  {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "success")
}
