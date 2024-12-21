package main

import (
	"fmt"
	"strings"
)

func main() {
	s2 := "how do you do"
	s3 := strings.Split(s2, " ")
	m1 := make(map[string]int, 10)
	for _, w := range s3 {
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
	}
	for key, value := range m1 {
		fmt.Println(key, value)
	}
}
