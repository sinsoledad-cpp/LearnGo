package models

import (
	"github.com/gin_qimi/bubble/dao"
)

/*
url		--> controller --> logic 	--> model
请求来了 --> 控制器 	--> 业务逻辑  --> 模型层的增删改查
*/

// Todo model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// Todo 增删改查
// CreateATodo 创建todo
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(todo).Error
	return
}

func GetAllTodo() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodo(id string) (todo *Todo, err error) {
	// todo = &Todo{}
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
