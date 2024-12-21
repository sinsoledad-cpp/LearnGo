package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		//goroutine 机制可以方便地实现异步处理
		//另外，在启动新的 goroutine 时，不应该使用原始上下文，必须使用它的只读副本
		copyContext := c.Copy()
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行:", copyContext.Request.URL.Path)
		}()
	})
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行:", c.Request.URL.Path)
	})
	r.Run(":8000")
}
