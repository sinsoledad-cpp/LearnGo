package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql" //init()
)

type user struct {
	Id   int
	Name string
	Age  int
}

// go连接mysql示例
var db *sqlx.DB //是一个连接池对象

func initDB() (err error) {
	//数据库信息
	//用户名：密码@tcp(ip:端口)/数据库的名字
	dsn := "root:sql3.1415@tcp(127.0.0.1:3307)/learn_go"
	//连接数据库
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10) //设置数据库连接池最大连接数
	db.SetMaxIdleConns(5)  //设置最大空闲连接数
	fmt.Println("链接数据库成功！")
	return
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
		return
	}
	sqlStr1 := `select id, name, age from user where id=1`
	var u user
	db.Get(&u, sqlStr1)
	fmt.Printf("u:%#v\n", u)

	var userList []user
	sqlStr2 := `select id,name,age from user`
	err = db.Select(&userList, sqlStr2)
	if err != nil {
		fmt.Printf("select failed,err:%v\n", err)
		return
	}
	fmt.Printf("useList:%#v\n", userList)
}
