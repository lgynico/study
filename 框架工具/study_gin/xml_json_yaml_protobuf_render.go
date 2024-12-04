package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lgynico/study-gin/testdata/pb"
)

func MessageTypeRender(e *gin.Engine) {

	e.GET("/someJSON", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	e.GET("moreJSON", func(ctx *gin.Context) {
		var msg struct {
			Name    string `json:"name"`
			Message string
			Number  int
		}

		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123

		ctx.JSON(http.StatusOK, msg)
	})

	e.GET("/someXML", func(ctx *gin.Context) {
		ctx.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	e.GET("/someYAML", func(ctx *gin.Context) {
		ctx.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	e.GET("/someProtoBuf", func(ctx *gin.Context) {
		ctx.ProtoBuf(http.StatusOK, &pb.Test{
			Label: "test",
			Reps:  []int64{1, 2},
		})
	})
}
