package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func PipeLine1() {
	tx := DB.TxPipeline()
	tx.Set("name", "fengfeng", 0)
	val := tx.Get("name")
	println(val.Val()) //nil 获取不到结果
	_, err := tx.Exec()
	if err != nil {
		panic(err) //事务失败
	}
	fmt.Println(DB.Get("name").Val())
	fmt.Println(val.Val())
}

func PipeLine2() {
	_, err := DB.Pipelined(func(tx redis.Pipeliner) error {
		tx.Set("price", 1001, 0)
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(DB.Get("price").Val())
}

func PipeLine3() {
	DB.Pipelined(func(tx redis.Pipeliner) error {
		tx.Set("price", 100, 0)
		val := tx.Get("price")
		fmt.Println(val.Val())
		_, err := tx.Exec()
		if err != nil {
			return err
		}
		fmt.Println(val.Val())
		return nil
	})
}
