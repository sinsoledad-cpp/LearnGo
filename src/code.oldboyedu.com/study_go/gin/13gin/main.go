package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 注册中间件
// 中间件要求是gin.HandlerFunc类型
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件执行了...")
		//设置变量到Context的key中，可以通过Get获取
		c.Set("request", "中间件")
		//执行中间件
		c.Next()
		//中间件执行完后续的处理
		status := c.Writer.Status() //返回状态码
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("耗时", t2)
	}
}

func main() {
	r := gin.Default()
	//注册中间件
	r.Use(MiddleWare()) //全局
	{
		r.GET("/MiddleWare", func(c *gin.Context) {
			//获取中间件设置的变量
			request, _ := c.Get("request")
			fmt.Println("中间件设置的变量:", request)
			//页面接收
			c.JSON(200, gin.H{"request": request})
		})
		r.GET("/MiddleWare2", MiddleWare(), func(c *gin.Context) {
			//获取中间件设置的变量
			request, _ := c.Get("request")
			fmt.Println("中间件设置的变量:", request)
			//页面接收
			c.JSON(200, gin.H{"request": request})
		})
		r.Run(":8000")
	}
}
