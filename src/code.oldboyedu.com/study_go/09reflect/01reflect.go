package main

import (
	"fmt"
	"reflect"
)

type myInt int64
// type name和type kind
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("type:%v\t,kind:%v\n", v.Name(), v.Kind())
}

func f1() {
	var a float32 = 3.14
	reflectType(a)
	var b int64 = 100
	reflectType(b)
}

func f2() {
	var a *float32
	var b myInt
	var c rune
	reflectType(a)
	reflectType(b)
	reflectType(c)

	type person struct {
		name string
		age  int
	}
	type book struct{ title string }
	var d = person{
		name: "沙河小王子",
		age:  10,
	}
	var e = book{title: "<a book>"}
	reflectType(d)
	reflectType(e)
}
// 通过反射获取值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	fmt.Printf("value:%v\n", v)
	fmt.Printf("type:%T\n", k)
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64,value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32,value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64,value is %f\n", float64(v.Float()))
	}
}

func f3() {
	var a float32 = 3.14
	var b int64 = 100
	reflectValue(a)
	reflectValue(b)

	c := reflect.ValueOf(10)
	fmt.Printf("type c:%T\n", c)
}

func main() {
	// f1()
	f2()
	// f3()
}
