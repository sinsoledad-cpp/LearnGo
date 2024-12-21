package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// redis
var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:8080",
		Password: "",
		DB:       0,
	})
	_, err = redisdb.Ping(ctx).Result()
	return
}

func redis01() {
	err := redisdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed,err:%v\n", err)
		return
	}
	val, err := redisdb.Get(ctx, "score").Result()
	if err != nil {
		fmt.Printf("get score failed,err:%v\n", err)
		return
	}
	fmt.Println("score", val)
	val2, err := redisdb.Get(ctx, "name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist!")
	} else if err != nil {
		fmt.Printf("get name failed,err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}

// 操作zset示例
func zsetDemo() {
	zsetkey := "language_rank"
	// 注意：v8版本使用[]*redis.Z；此处为v9版本使用[]redis.Z
	languages := []redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	ctxs, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	err := redisdb.ZAdd(ctxs, zsetkey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed,err:%v\n", err)
		return
	}
	fmt.Println("zadd success")

	newScore, err := redisdb.ZIncrBy(ctxs, zsetkey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed,err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	ret := redisdb.ZRevRangeWithScores(ctxs, zsetkey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = redisdb.ZRangeByScoreWithScores(ctxs, zsetkey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed, err:%v\n", err)
		return
	}
	fmt.Println("connect redis success!")
	// redis01()
	zsetDemo()
}
