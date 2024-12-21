package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		//表单取文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		//传到项目根目录，名字就用本身的
		c.SaveUploadedFile(file, file.Filename)
		//打印信息
		c.String(200, fmt.Sprintf("%s upload!", file.Filename))
	})
	r.Run(":8000")
}
