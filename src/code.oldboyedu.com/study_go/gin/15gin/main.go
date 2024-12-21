package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "no set"
			//给客户端设置cookie
			//name string, cookie的名字
			//value string, cookie的值
			//maxAge int, cookie的过期时间
			//path string, cookie的路径,cookie所在的目录
			//domain string,cookie的作用域 ,域名
			//secure bool, 是否是安全的,是否只能同通过https访问
			//httpOnly bool,是否是httpOnly,是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/", "localhost", false, true)
		}
		fmt.Println("cookie:", cookie)

	})
	r.Run(":8000")
}
