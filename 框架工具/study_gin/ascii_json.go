package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AsciiJSON(e *gin.Engine) {
	e.GET("/asciiJSON", func(ctx *gin.Context) {
		data := map[string]any{
			"lang": "GO 语言",
			"tag":  "<br>",
		}

		// 使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON
		ctx.AsciiJSON(http.StatusOK, data)
	})
}
