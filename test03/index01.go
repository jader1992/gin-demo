package main

import "github.com/gin-gonic/gin"

type LoginForm struct {
	User string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	route := gin.Default()
	route.POST("/login", func(c *gin.Context) {
		var form LoginForm
		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logging in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})
	route.Run(":8080")
}
