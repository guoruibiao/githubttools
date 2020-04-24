package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guoruibiao/githubtools/resume"
	"github.com/guoruibiao/githubtools/timemap"
	"net/http"
	"strings"
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
	
	app.GET("/location", func(context *gin.Context) {
		url := context.Query("url")
		profileParser := resume.ProfileParser{}
		if strings.Contains(url, "issue") {
			issueParser := resume.IssueParser{}
			issueParser.Parse(url)
			profileParser.Homepage = issueParser.Homepage
			profileParser.Username = issueParser.Username
		}else if strings.Contains(url, "issue") == false && strings.Contains(url, "github.com/") {
			if len(url) > 11 { // contains github username
				arr := strings.Split(url, "/")
				profileParser.Username = arr[1]
				profileParser.Homepage = "https://" + strings.Join(arr[:2], "/")
			}else{
				context.JSON(http.StatusOK, "参数格式非法")
			}
		}else{
			// 当做是 username 处理
			profileParser.Username = url
			profileParser.Homepage = "https://github.com/" + url
		}
		
		profileParser.Parse("li>span.p-label")
		context.JSON(http.StatusOK, profileParser)
	})
	
	app.Run(":9999")
}
