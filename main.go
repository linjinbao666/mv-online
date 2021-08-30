package main

import (
	"mv-online/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func startWeb() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "首页",
		})
	})
	router.GET("/video", func(c *gin.Context) {
		c.HTML(http.StatusOK, "video.tmpl", gin.H{
			"title": "视频列表",
		})
	})
	router.GET("/api/video/list", func(c *gin.Context) {
		videos := pkg.Videos("", "", "")
		data := gin.H{"data": videos, "code": 0, "msg": "", "count": 10}
		c.JSON(http.StatusOK, data)
	})

	router.Static("/static/video", "./videos")
	router.Run()
}

func main() {
	startWeb()
}
