package main

// 结构体与json
// 1.序列化:把Go语言中的结构体变量-->json格式的字符串
// 2.反序列化:json格式的字符串--> Go语言中能够识别的结构体变量
import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "周琳",
		Age:  19,
	}
	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
		return
	}
	fmt.Printf("%v\n", string(b))

	b,err=json.MarshalIndent(p1,"","\t")
	if err!=nil{
		panic(err)
	}
	fmt.Println(string(b))

	// 反序列化
	str := `{"name":"理想","age":18}`
	fmt.Printf("%T\n", str)
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%#v\n", p2)
}
