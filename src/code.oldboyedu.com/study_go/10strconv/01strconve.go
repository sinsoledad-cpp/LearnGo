package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "10000"
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("parseint failed, err: ", err)
		return
	}
	fmt.Printf("%#v %T\n", ret1, ret1)

	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n", retInt, retInt)

	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue)

	floatStr := "1.234"
	floatValue, _ := strconv.ParseFloat(floatStr,64)
	fmt.Printf("%#v %T\n", floatValue, floatValue)

	i:=97
	ret2:=string(i)
	fmt.Printf("%#v %T\n", ret2, ret2)
	ret3:=fmt.Sprintf("%d",i)
	fmt.Printf("%#v %T\n", ret3, ret3)
	ret4:=strconv.Itoa(i)
	fmt.Printf("%#v %T\n", ret4, ret4)

}
