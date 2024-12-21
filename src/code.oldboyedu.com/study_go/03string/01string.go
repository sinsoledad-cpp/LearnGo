package main

import (
	"fmt"
	"unicode"
)

func main() {
	s1 := "heofioasjfo你似乎还是"
	var count int
	for _, c := range s1 {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println(count)
}