package main

import (
	"github.com/CloudWeGo/gomall/demo/demo_proto/biz/dal"
	"github.com/CloudWeGo/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/CloudWeGo/gomall/demo/demo_proto/biz/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()

	//CURD

	// mysql.DB.Create(&model.User{Email: "test@gmail.com", Password: "123456"})

	// mysql.DB.Model(&model.User{}).Where("email = ?", "test@gmail.com").Update("password", "1234512322")

	// var row model.User
	// mysql.DB.Model(&model.User{}).Where("email = ?", "test@gmail.com").First(&row)
	// fmt.Printf("%+v\n", row)

	// mysql.DB.Where("email = ?", "test@gmail.com").Delete(&model.User{})

	mysql.DB.Unscoped().Where("email = ?", "test@gmail.com").Delete(&model.User{})
}
