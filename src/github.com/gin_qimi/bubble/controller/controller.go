package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin_qimi/bubble/models"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 会发请求到这里
	// 1.从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	// 2.存入数据库
	// err := DB.Create(&todo).Error
	// if err != nil {
	// }
	if err := models.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
		// c.JSON(http.StatusOK, gin.H{
		// 	"code": 2000,
		// 	"msg":  "success",
		// 	"data": todo,
		// })
	}
	// 3.反回响应
}

func GetTodoList(c *gin.Context) {
	//查询todo这个表中的所有数据
	if todoList, err := models.GetAllTodo(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
	c.BindJSON(&todo) //从请求中把数据拿出来
	if err = models.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			id: "deleted",
		})
	}
}
