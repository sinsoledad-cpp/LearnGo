package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取客户端cookie并验证
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"status": "401"})
		c.Abort()
		return
	}
}
func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		//设置cookie
		c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
		//返回信息
		c.String(200, "login success")
	})
	r.GET("/home", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "200"})
	})
	r.Run(":8000")
}
