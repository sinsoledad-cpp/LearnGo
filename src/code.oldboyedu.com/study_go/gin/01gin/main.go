package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getting(c *gin.Context) {

}

func f1() {
	//1.创建路由
	r := gin.Default() //默认使用了两个中间件：Logger（），Recovery（）
	//r:=gin.New()//创建不带中间件的路由

	//2.绑定路由规制，执行的函数
	//gin.Context,封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})
	r.POST("/xxxPost", getting)
	r.PUT("/xxxPut")
	//3.监听端口，默认在8080
	r.Run(":8000")
}

func f2() {
	r := gin.Default()
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name+" + "+action)
	})
	r.Run(":8000")
}

func f3() {
	r := gin.Default()
	r.GET("welcome", func(c *gin.Context) {
		//defaultQuery第二个参数的默认值
		//localhost/welcome
		//localhost/welcome?name=hhh
		name := c.DefaultQuery("name", "Jack")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	r.Run(":8000")
}

func main() {
	// f2()
	f3()
}
