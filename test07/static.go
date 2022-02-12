package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	route.Static("/assert", "./assets")
	route.StaticFS("/more_static", http.Dir("my_file_system"))
	route.StaticFile("/favicon.ico", "./resources/favicon.ico")

	route.Run(":8080")
}
