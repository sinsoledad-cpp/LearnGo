package main

import (
	"fmt"
	"reflect"
)

// 通过反射设置变量的值
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
func f1() {
	var a int64 = 100
	// reflectSetValue1(a)
	reflectSetValue2(&a)
	fmt.Println(a)
}

func main() {
	f1()
}
