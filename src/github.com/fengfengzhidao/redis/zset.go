package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func RedisZset() {
	DB.ZAdd("z",
		redis.Z{Score: 45, Member: "fengfeng0"},
		redis.Z{Score: 35, Member: "fengfeng2"},
		redis.Z{Score: 40, Member: "fengfeng1"})

	fmt.Println(DB.ZRange("z", 0, -1).Val())    //升序
	fmt.Println(DB.ZRevRange("z", 0, -1).Val()) //降序

	fmt.Println(DB.ZRangeByScore("z", redis.ZRangeBy{Min: "40", Max: "50"}).Val())
	fmt.Println(DB.ZRevRangeByScore("z", redis.ZRangeBy{Min: "40", Max: "50"}).Val())

	fmt.Println(DB.ZRangeWithScores("z", 0, -1).Val())
}
