package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 使用 SecureJSON 防止 json 劫持。如果给定的结构是数组值，则默认预置 "while(1)," 到响应体。
func SecureJSON(e *gin.Engine) {

	// 你也可以使用自己的 SecureJSON 前缀
	// r.SecureJsonPrefix(")]}',\n")
	e.GET("/securejson", func(ctx *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		// output: while(1);["lena","austin","foo"]
		ctx.SecureJSON(http.StatusOK, names)
	})
}
