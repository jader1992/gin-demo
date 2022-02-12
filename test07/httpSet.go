package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 方法1：
	route := gin.Default()
	http.ListenAndServe(":8080", route)

	// 方法2：
	s := &http.Server{
		Addr:              ":8080",
		Handler:           nil,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	s.ListenAndServe()
}
