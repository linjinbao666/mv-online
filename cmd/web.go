package cmd

import (
	"embed"
	"html/template"
	"mv-online/pkg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

//go:embed templates/*
var f embed.FS

func startWeb(addr string, workingDir string) {
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
		videos := pkg.Videos("", "", "", workingDir)
		data := gin.H{"data": videos, "code": 0, "msg": "", "count": 10}
		c.JSON(http.StatusOK, data)
	})

	router.GET("/dashboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{
			"title": "面板",
		})

	})

	router.Static("/static/video", workingDir+"/videos")
	router.Run(addr)
}

var port int
var workingDir string

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "启动web服务",
	Long:  `启动web服务`,
	Run: func(cmd *cobra.Command, args []string) {
		addr := ":" + strconv.Itoa(port)
		startWeb(addr, workingDir)
	},
}

func init() {
	rootCmd.AddCommand(webCmd)

	webCmd.Flags().IntVarP(&port, "port", "p", 8080, "--port=8080")
	webCmd.MarkFlagRequired("port")

	webCmd.Flags().StringVarP(&workingDir, "data", "", "", "--data=/opt/mv-online")
	webCmd.MarkFlagRequired("data")

}
