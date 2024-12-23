package main

import "fmt"

func RedisHash() {
	DB.HSet("info", "name", "fengfeng")
	DB.HSet("info", "age", 23)

	fmt.Println(DB.HGet("info", "name"))
	fmt.Println(DB.HGet("info", "age").Val())

	fmt.Println(DB.HGetAll("info").Val())
	fmt.Println(DB.HKeys("info").Val())
	fmt.Println(DB.HVals("info").Val())
	fmt.Println(DB.HLen("info").Val())
}
