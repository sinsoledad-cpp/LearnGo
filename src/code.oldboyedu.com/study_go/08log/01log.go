package main

import (
	"fmt"
	"log"
	"os"
)

func f1() {
	log.Println("this is a normal log")
	v := "normal"
	log.Printf("this is a %v log\n", v)
	log.Fatalln("this is a fatal log")
	log.Panicln("this is a panic log")
}
func f2() {
	logFile, err := os.OpenFile("./log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed,err: ", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("this is a normal log")
	log.SetPrefix("[hello world]")
	log.Println("this is a normal log")
}
func f3() {
	logger:=log.New(os.Stdout,"<new>",log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("this is a new log")

}
func main() {
	// f1()
	// f2()
	f3()
}
