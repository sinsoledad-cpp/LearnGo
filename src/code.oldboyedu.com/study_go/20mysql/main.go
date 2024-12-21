package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //init()
)

// 定义一个全局对象db
var db *sql.DB

type user struct {
	id   int
	age  int
	name string
}

// 定义一个初始化数据库的函数
func initDB() (err error) {
	//数据库信息
	dsn := "root:sql3.1415@tcp(127.0.0.1:3307)/learn_go"
	//链接数据库
	db, err = sql.Open("mysql", dsn) //不会校验用户名和密码时候错误
	if err != nil {                  //dsn格式不正确的时候会报错
		fmt.Printf("dsn:%s invalid,err:%v\n", dsn, err)
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed,err:%v\n", dsn, err)
		return err
	}
	db.SetMaxOpenConns(10) //设置数据库连接池最大连接数
	db.SetMaxIdleConns(5)  //设置最大空闲连接数
	fmt.Println("链接数据库成功！")
	return nil
}

// 查询单个记录
func queryOne(id int) {
	var u user
	//sql语句
	sqlStr := `select id,name,age from user where id=?`
	// //执行
	// rowObj := db.QueryRow(sqlStr, id)
	// //拿到结果
	// rowObj.Scan(&u.id, &u.name, &u.age)
	//执行并拿到结果
	db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)
	fmt.Printf("u:%#v\n", u)
}

// 查询多个记录
func queryMore(id int) {
	//sql语句
	sqlStr := `select id,name,age from user where id>?`
	//执行
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("exec %s query failed, err:%v\n", sqlStr, err)
		return
	}
	//一定要关闭rows
	defer rows.Close()
	//循环取值
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("u:%#v\n", u)
	}
}
func insert() {
	//写sql语句
	sqlStr := `insert into user(name,age) values("朝阳",20)`
	//exec
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	//如果插入数据,能够拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Println("id:\n", id)
}

// 更新操作
func updateRow(newAge int, id int) {
	sqlStr := `update user set age=? where id>?`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("update failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Printf("更新了%d行数据\n", n)
}

// 删除
func deleteRow(id int) {
	sqlStr := `delete from user where id=?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Printf("删除了%d行数据\n", n)
}

// 预处理方式插入多条数据
func prepareInsert() {
	sqlStr := `insert into user(name,age) values(?,?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
		return
	}
	defer stmt.Close()
	//后续只需要拿到stmt去执行一些操作
	var m = map[string]int{
		"六七强": 30,
		"烷相机": 32,
		"天说":  72,
		"百惠":  40,
	}
	for k, v := range m {
		stmt.Exec(k, v) //后续只需要传值
	}
}

// 事务操作示例
func transactionDemo() {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set age=age+1 where id=?"
	ret1, err := tx.Exec(sqlStr1, 1)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := "Update user set age=age-1 where id=?"
	ret2, err := tx.Exec(sqlStr2, 2)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	fmt.Println(affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交啦...")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		fmt.Println("事务回滚啦...")
	}

	fmt.Println("exec trans success!")
}

// go链接mysql
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	// queryOne(1)
	// insert()
	// queryMore(2)
	// deleteRow(4)
	// prepareInsert()
	// queryMore(2)
	// updateRow(18, 2)
	transactionDemo()
}
