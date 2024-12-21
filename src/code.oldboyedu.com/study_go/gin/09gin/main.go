package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//1.json
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "someJSON", "status": 200})
	})
	//2.结构体响应
	r.GET("someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(200, msg)
	})
	//3.mxl
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "someXML", "status": 200})
	})
	//4.yaml响应
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{"message": "someYaml", "status": 200})
	})
	//5.protobuf响应,谷歌开发的高校储存读取工具
	//数组？切片？如果自己构建一个传输格式,应该是什么格式？
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		//定义数据
		label := "label"
		//传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})
	r.Run(":8000")
}
