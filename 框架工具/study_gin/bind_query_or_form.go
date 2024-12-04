package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// curl -X GET "localhost:8080/bindQueryOrForm?name=appleboy&address=xyz&birthday=1992-03-15"
func BindQueryOrForm(e *gin.Engine) {
	type person struct {
		Name     string    `form:"name"`
		Address  string    `form:"address"`
		Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	}

	// 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
	// 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）。
	e.GET("/bindQueryOrForm", func(ctx *gin.Context) {
		var person person
		if ctx.ShouldBind(&person) == nil {
			log.Println(person.Name)
			log.Println(person.Address)
			log.Println(person.Birthday)
		}
		ctx.String(200, "Success")
	})
}
