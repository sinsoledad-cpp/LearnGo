package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	// binging:"required" 修饰字段，如果接收为空值，则报错，是必须字段
	User     string `form:"username" json:"user" uri:"user" xml:"user" binging:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binging:"required"`
}

// Form数据格式
func main() {
	r := gin.Default()
	//Form绑定
	r.POST("/loginForm", func(c *gin.Context) {
		//声明接收的变量
		var from Login
		if err := c.Bind(&from); err != nil {
			//返回错误信息
			//gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//判断用户名密码是否正确
		if from.User != "root" || from.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run(":8000")
}
