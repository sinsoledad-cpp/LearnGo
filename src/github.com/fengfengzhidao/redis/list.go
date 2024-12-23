package main

import "fmt"

func RedisList() {
	DB.RPush("list", "wangwu", "zhangsan", "lisi")
	DB.LPush("list", "wangwu1", "zhangsan2", "lisi3")
	fmt.Println(DB.LRange("list", 0, -1))
	fmt.Println(DB.LRange("list", 0, 0).Val())

	fmt.Println(DB.LPop("list"))
	fmt.Println(DB.RPop("list"))
}
