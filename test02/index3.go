package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
	route := gin.Default()
	// 定义分隔符
	route.Delims("{[{", "}]}")
	// 定义自定义函数
	route.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})

	// 导入模版文件
	route.LoadHTMLFiles("./templates/testdata/raw.tmpl")

	route.GET("/raw", func(context *gin.Context) {
		context.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
			"now": time.Date(2021, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})

	route.Run(":8080")


}

