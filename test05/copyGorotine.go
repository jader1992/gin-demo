package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main()  {
	// 当在中间件或 handler 中启动新的 Goroutine 时，不能使用原始的上下文，必须使用只读副本。
	route := gin.Default()

	route.GET("/long_async", func(context *gin.Context) {
		// 创建在 goroutine 中使用的副本
		cp := context.Copy()
		go func() {
			time.Sleep(5 * time.Second)

			// 请注意您使用的是复制的上下文 "cCp"，这一点很重要
			log.Println("Done ! in path" + cp.Request.URL.Path)
		}()
	})

	route.GET("/long_sync", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done ! in path" + context.Request.URL.Path)
	})

	route.Run(":8080")
}
