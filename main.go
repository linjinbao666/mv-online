package main

import (
	"embed"
	"html/template"
	"mv-online/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var f embed.FS

func startWeb() {
	router := gin.Default()
	templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl"))
	router.SetHTMLTemplate(templ)
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
