package main

import (
	"fmt"
	"time"
)

func RedisString() {
	DB.Set("name", "fengfeng", 0) //0 永不过期
	DB.Set("age", 23, 0)          //0 永不过期
	// status := DB.Set("name", "fengfeng", 0) //0 永不过期
	// fmt.Println(status.Result())
	// fmt.Println(status.Val())
	// fmt.Println(status.Err())

	// name, err := DB.Get("name").Result()

	stringCmd := DB.Get("name")
	fmt.Println(stringCmd.Result())
	fmt.Println(stringCmd.Val()) //string类型用val
	fmt.Println(stringCmd.Err())

	stringCmd = DB.Get("age")
	fmt.Println(stringCmd.Result())
	fmt.Println(stringCmd.Int())   //int类型用int
	fmt.Println(stringCmd.Int64()) //int类型用int64
	fmt.Println(stringCmd.Err())

	fmt.Println(DB.Exists("name").Val())  //判断key是否存在 1存在 0不存在
	fmt.Println(DB.Exists("name1").Val()) //判断key是否存在 1存在 0不存在

	fmt.Println(DB.Incr("age").Val())      //自增1
	fmt.Println(DB.IncrBy("age", 2).Val()) //自增value
	fmt.Println(DB.Decr("age").Val())      //自减1
	fmt.Println(DB.DecrBy("age", 2).Val()) //自减value
	fmt.Println(DB.Del("age").Val())       //删除key

	fmt.Println(DB.TTL("name"))               //-1s 永不过期
	DB.Set("name", "fengfeng", 2*time.Second) //
	fmt.Println(DB.TTL("name"))               // 查看key的过期时间 2s
	time.Sleep(3 * time.Second)
	fmt.Println(DB.TTL("name")) //-2s 已过期

	DB.Set("name", "fengfeng", 0)
	DB.Expire("name", 20*time.Second)
}
