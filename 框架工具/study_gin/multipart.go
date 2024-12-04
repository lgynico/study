package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func MultipartBind(e *gin.Engine) {
	e.POST("/login", func(ctx *gin.Context) {
		var form LoginForm
		if ctx.ShouldBind(&form) == nil {
			if form.User == "root" && form.Password == "123456" {
				ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})
}
