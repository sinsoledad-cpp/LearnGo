package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var DB *redis.Client

func RedisConnect() {
	client := redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		Password:    "", // no password set
		DB:          15, // use default DB
		DialTimeout: time.Second,
	})
	err := client.Ping().Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis connect success")
	DB = client
}
