package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/", func(context *gin.Context) {
		if pusher := context.Writer.Pusher(); pusher != nil {
			if err := pusher.Push("/asserts/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		context.HTML(200, "https", gin.H{
			"status": "success",
		})
	})

	r.RunTLS(":8080", "./testdata/server/pem", "./testdata/server.key") // 启动https服务
}