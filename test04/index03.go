package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	var imgUrl string = "http://img.netbian.com/file/2021/1008/049d8ebda647bb8df7f037a17791d253.jpg"

	router.GET("/someDataFromReader", func(context *gin.Context) {
		response, err := http.Get(imgUrl)
		if err != nil || response.StatusCode != http.StatusOK {
			context.Status(http.StatusServiceUnavailable)
			return
		}

		if response == nil {
			context.String(http.StatusOK, fmt.Sprintf("image can not open"))
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-type")

		extraHeaders := map[string]string{
			"content-Disposition": `attachment; filename="gophper.png"`,
		}

		// 访问 /dataFromReader 路由会下载图片到本地。
		context.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	router.Run(":8080")
}
