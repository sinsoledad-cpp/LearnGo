package main

import (
	"flag"
	"fmt"
	"time"
)

func f1() {
	//创建一个标志位参数
	name := flag.String("name", "王子", "请输入名字")
	age := flag.Int("age", 19, "年龄")
	married := flag.Bool("married", false, "are you haved married?")
	cTime := flag.Duration("ct", time.Second, "多久了")
	flag.Parse() //使用flag
	fmt.Println(name)
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*cTime)
	fmt.Printf("%T\n", *cTime)

	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())
}

func f2() {
	var name string
	flag.StringVar(&name, "name", "王治", "请输入名字")
	flag.Parse()
	fmt.Println(name)
}

// flag获取命令行参数
func main() {
	f1()
	// f2()
}
