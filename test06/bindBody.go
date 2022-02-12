package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

//c.ShouldBindBodyWith 会在绑定之前将 body 存储到上下文中。 这会对性能造成轻微影响，如果调用一次就能完成绑定的话，那就不要用这个方法。
//只有某些格式需要此功能，如 JSON, XML, MsgPack, ProtoBuf。 对于其他格式, 如 Query, Form, FormPost, FormMultipart 可以多次调用 c.ShouldBind() 而不会造成任任何性能损失

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func someHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}

	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		c.String(http.StatusOK, "test")
	}
}

func SomeHandlerMulti(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// 读取 c.Request.Body 并将结果存入上下文。
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// 这时, 复用存储在上下文中的 body。
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, `the body should be formB JSON`)
		// 可以接受其他格式
	} else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
		c.String(http.StatusOK, `the body should be formB XML`)
	} else {
		c.String(http.StatusOK, "test")
	}
}

func main() {
	route := gin.Default()

	route.POST("/testingSingle", someHandler)
	route.POST("/testingMulti", SomeHandlerMulti)

	route.Run(":8080")
}
