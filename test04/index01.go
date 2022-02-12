package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusOK, fmt.Sprintf("文件上传失败 '%s'", err))
		}

		filename := file.Filename
		log.Println(filename)
		log.Println("./dst/"+filename)
		// 上传文件至指定目录
		c.SaveUploadedFile(file, "./dst/"+filename)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8080")
}
