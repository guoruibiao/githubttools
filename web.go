package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guoruibiao/githubtools/timemap"
	"net/http"
)

func main() {
	app := gin.Default()
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "It works.")
	})
	
	app.GET("/timemap", func(context *gin.Context) {
		cityname := context.Query("cityname")
		mapper := timemap.NewTimeMap(cityname)
		if err := mapper.DoQuery(); err != nil {
			context.JSON(http.StatusOK, err.Error())
		}else {
			context.JSON(http.StatusOK, mapper.GetCityTimes())
		}
	})
	
	app.Run(":9999")
}
