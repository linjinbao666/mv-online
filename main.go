package main

import (
	"mv-online/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func startWeb() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/video", func(c *gin.Context) {
		c.HTML(http.StatusOK, "video.tmpl", gin.H{
			"title": "视频列表",
		})
	})
	router.GET("/api/video", func(c *gin.Context) {
		videos := pkg.Videos("", "", "")
		data := gin.H{"data": videos, "code": 0, "msg": "", "count": 10}
		c.JSON(http.StatusOK, data)
	})
	router.Run()
}

func main() {
	startWeb()
}
