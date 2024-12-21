package main

import (
	"fmt"

	"code.oldboyedu.com/study_go/17test/01test/split_string"
)

func main() {
	ret := split_string.Split("babcbef","b")
	fmt.Printf("%#v\n",ret)
}