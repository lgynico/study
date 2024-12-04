package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

/*
POST /post?id=1234&page=1 HTTP/1.1
Content-Type: application/x-www-form-urlencoded

name=manu&message=this_is_great
*/
func QueryAndPostForm(e *gin.Engine) {
	e.POST("/post", func(ctx *gin.Context) {
		var (
			id      = ctx.Query("id")
			page    = ctx.DefaultQuery("page", "0")
			name    = ctx.PostForm("name")
			message = ctx.PostForm("message")
		)

		log.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)

	})
}
