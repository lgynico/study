package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Form(e *gin.Engine) {
	e.POST("/form_post", func(ctx *gin.Context) {
		var (
			message = ctx.PostForm("message")
			nick    = ctx.DefaultPostForm("nick", "anonymous")
		)

		ctx.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
}
