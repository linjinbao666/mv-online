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

	router.DELETE("/api/video/:name", func(c *gin.Context) {
		name := c.Param("name")
		_, err := pkg.VideoDelete(name, "", "", workingDir)
		if err != nil {
			data := gin.H{"data": name, "code": -1, "msg": "删除失败" + err.Error()}
			c.JSON(http.StatusOK, data)
		}
		c.JSON(http.StatusOK, gin.H{
			"data": name, "code": 0, "msg": "删除成功！",
		})
	})

	router.GET("/dashboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{
			"title": "面板",
		})

	})

	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		var videoFolder = workingDir + "/videos"
		c.SaveUploadedFile(file, videoFolder+"/"+file.Filename)
		data := gin.H{"data": file.Filename, "code": 0, "msg": "上传成功"}
		c.JSON(http.StatusOK, data)
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
