package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func myTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	since := time.Since(start)
	c.String(200, since.String())
	fmt.Println("耗时:", since)
}

func main() {
	r := gin.Default()
	r.Use(myTime)
	shoppingGroup := r.Group("/shopping")
	{
		//购物车
		shoppingGroup.GET("/cart", func(c *gin.Context) {
			time.Sleep(2 * time.Second)
			c.String(200, "购物车")
		})
		//订单
		shoppingGroup.GET("/order", func(c *gin.Context) {
			time.Sleep(3 * time.Second)
			c.String(200, "订单")
		})
	}
	r.Run(":8000")
}
