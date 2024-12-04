package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONP(e *gin.Engine) {
	e.GET("/JSONP", func(ctx *gin.Context) {
		data := map[string]any{
			"foo": "bar",
		}

		// /JSONP?callback=x
		// output：x({"foo":"bar"})
		ctx.JSONP(http.StatusOK, data)
	})
}
