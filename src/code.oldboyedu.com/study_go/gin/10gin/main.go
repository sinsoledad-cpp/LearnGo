package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	// r.LoadHTMLFiles("templates/index.tmpl")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "我的标题",
		})
	})
	r.Run(":8000")
}
