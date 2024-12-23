package main

import "fmt"

func RedisSet() {
	DB.SAdd("set", 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(DB.SMembers("set").Val())    //遍历元素，按序处理
	fmt.Println(DB.SMembersMap("set").Val()) //高效检查元素是否存在，不关心顺序
	fmt.Println(DB.SIsMember("set", 4).Val())
	fmt.Println(DB.SIsMember("set", 8).Val())

	DB.SAdd("set1", 1, 2, 3)
	DB.SAdd("set2", 2, 3, 4)

	fmt.Println(DB.SInter("set1", "set2").Val())
	fmt.Println(DB.SUnion("set1", "set2").Val())

	fmt.Println(DB.SDiff("set1", "set2").Val())
	fmt.Println(DB.SDiff("set2", "set1").Val())

	fmt.Println(DB.SPop("set"))
	fmt.Println(DB.SPopN("set", 2))
}
