package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func RedisWatch() {
	DB.Watch(func(tx *redis.Tx) error {
		_, err := tx.Pipelined(func(pipe redis.Pipeliner) error {
			time.Sleep(time.Second * 5)
			pipe.Set("price", 100, 0)
			return nil
		})
		if err != nil {
			return err
		}
		fmt.Println(tx.Get("price").Val())
		return nil
	}, "price")
}
