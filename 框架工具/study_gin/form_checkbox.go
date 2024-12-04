package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FormCheckbox(e *gin.Engine) {
	type myForm struct {
		Colors []string `form:"colors[]"`
	}

	e.POST("/formCheckbox", func(ctx *gin.Context) {
		var fakeForm myForm
		_ = ctx.ShouldBind(&fakeForm)
		ctx.JSON(http.StatusOK, gin.H{"color": fakeForm.Colors})
	})
}
