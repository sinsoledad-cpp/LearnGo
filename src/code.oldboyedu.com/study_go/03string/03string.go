package main

import "fmt"

func main() {
	ss := "a阿佛阿a"
	r := make([]rune, 0, len(ss))
	for _, c := range ss {
		r = append(r, c)
	}
	fmt.Println("[]rune: ", r)
	for i:=0;i<len(r)/2;i++{
		if r[i]!=r[len(r)-1-i]{
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}
