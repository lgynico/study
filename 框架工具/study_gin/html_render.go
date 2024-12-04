package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HTMLRender(e *gin.Engine) {
	// e.LoadHTMLGlob("templates/*")
	e.LoadHTMLGlob("templates/**/*")

	e.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	e.GET("/posts/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})

	e.GET("/users/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
}

func Must(e *gin.Engine) {
	html := template.Must(template.ParseFiles("file1", "file2"))
	e.Delims("{[{", "}]}")
	e.SetHTMLTemplate(html)
}

func LoadHTMLFiles(e *gin.Engine) {
	e.Delims("{[{", "}]}")
	e.SetFuncMap(template.FuncMap{
		"formatAsDate": formateAsDate,
	})
	e.LoadHTMLFiles("./testdata/template/raw.tmpl")

	e.GET("/raw", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "raw.tmpl", map[string]any{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})
}

func formateAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}
