package main

import "github.com/gin-gonic/gin"

// 使用中间件

func main() {
	r := gin.New()

	// 添加全局中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 也可以为路由添加任意数量的中间件
	r.GET("/benchmark", MyBenchLogger(), benchEndpoint) // 此处的MyBenchLogger()是中间件，benchEndpoint是处理函数

	// 认证路由
	// 方法1： authorized := r.Group("/", AuthRequired())
	// 方法2：
	authorized := r.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)
	}

	// 嵌套路由组
	testing := authorized.Group("testing")
	testing.GET("/analytics", analyticsEndpoint)

	r.Run(":8080")
}


