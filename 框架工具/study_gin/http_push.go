package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HTTPPush(e *gin.Engine) {
	e.Static("/assets", "./assets")
	e.SetHTMLTemplate(template.Must(template.New("https").Parse(`
<html>
<head>
	<title>Https Test</title>
	<script src="/assets/app.js"></script>
</head>
<body>
	<h1 style="color: red">Welcome, Ginner!</h1>
</body>
</html>
	`)))

	e.GET("/", func(ctx *gin.Context) {
		if pusher := ctx.Writer.Pusher(); pusher != nil {
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}

		ctx.HTML(http.StatusOK, "https", gin.H{
			"status": "success",
		})
	})
}
