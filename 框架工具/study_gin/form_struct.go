package main

import "github.com/gin-gonic/gin"

func FormStruct(e *gin.Engine) {

	type (
		structA struct {
			A string `form:"a"`
		}
		structB struct {
			NestedStruct structA
			B            string `form:"b"`
		}
		structC struct {
			NestedStructPointer *structA
			C                   string `form:"c"`
		}
		structD struct {
			NestedAnonyStruct struct {
				X string `form:"x"`
			}
			D string `form:"d"`
		}
	)

	e.GET("/getb", func(ctx *gin.Context) {
		var b structB
		ctx.Bind(&b)
		ctx.JSON(200, gin.H{
			"a": b.NestedStruct,
			"b": b.B,
		})
	})
	e.GET("/getc", func(ctx *gin.Context) {
		var c structC
		ctx.Bind(&c)
		ctx.JSON(200, gin.H{
			"a": c.NestedStructPointer,
			"c": c.C,
		})
	})
	e.GET("/getd", func(ctx *gin.Context) {
		var d structD
		ctx.Bind(&d)
		ctx.JSON(200, gin.H{
			"x": d.NestedAnonyStruct,
			"d": d.D,
		})
	})
}
