package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	//数据库信息
	//用户名：密码@tcp(ip:端口)/数据库的名字
	dsn := "root:sql3.1415@tcp(127.0.0.1:3307)/study_mysql"
	//连接数据库
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	db.SetMaxIdleConns(16)  //设置最大空闲连接数
	fmt.Println("链接数据库成功！")
	return
}

func queryAllBook() (bookList []*Book, err error) {
	sqlStr := "select id,title,price from book"
	err = db.Select(&bookList, sqlStr)
	if err != nil {
		fmt.Println("query default")
		return
	}
	return
}

func insertBook(title string, price int64) (err error) {
	sqlStr := "insert into book(title,price) values(?,?)"
	_, err = db.Exec(sqlStr, title, price)
	if err != nil {
		fmt.Println("insert default")
		return
	}
	return
}

func deleteBook(id int64) (err error) {
	sqlStr := "delete from book where id=?"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("delete default")
		return
	}
	return
}
