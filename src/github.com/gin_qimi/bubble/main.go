package main

import (
	"github.com/gin_qimi/bubble/dao"
	"github.com/gin_qimi/bubble/models"
	"github.com/gin_qimi/bubble/routers"
)

func main() {
	//创建数据库
	//sql:CREATE DATABASE bubble;

	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close() //程序退出关闭数据库连接

	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	//注册路由
	r := routers.SetupRouter()
	r.Run()
}
