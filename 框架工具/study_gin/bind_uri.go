package main

import "github.com/gin-gonic/gin"

type person struct {
	ID   string `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func BindURI(e *gin.Engine) {

	e.GET("/:name/:id", func(ctx *gin.Context) {
		var person person
		if err := ctx.ShouldBindUri(&person); err != nil {
			ctx.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})
}
