package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PureJSON(e *gin.Engine) {
	e.GET("/json", func(ctx *gin.Context) {
		// 提供 unicode 实体
		ctx.JSON(http.StatusOK, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	e.GET("/purejson", func(ctx *gin.Context) {
		// 提供字面字符
		ctx.PureJSON(http.StatusOK, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
}
