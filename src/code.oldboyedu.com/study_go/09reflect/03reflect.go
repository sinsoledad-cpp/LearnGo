package main

import (
	"fmt"
	"reflect"
)

func f() {
	//IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())

	b := struct{}{} //实例化一个匿名结构体
	fmt.Println("no exist member:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	fmt.Println("no exist way:", reflect.ValueOf(b).MethodByName("abc").IsValid())

	c := map[string]int{}
	fmt.Println("map no exist key:", reflect.ValueOf(c).MapIndex(reflect.ValueOf("ok:")).IsValid())
}

func main() {
	f()
}
