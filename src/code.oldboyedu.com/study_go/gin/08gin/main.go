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

// json数据格式
func main() {
	r := gin.Default()
	//json绑定
	//http://localhost/root/admin
	r.GET("/:user/:password", func(c *gin.Context) {
		//声明接收的变量
		var login Login
		if err := c.ShouldBindUri(&login); err != nil {
			//返回错误信息
			//gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//判断用户名密码是否正确
		if login.User != "root" || login.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run(":8000")
}
