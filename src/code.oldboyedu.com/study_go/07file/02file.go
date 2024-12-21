package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeFromFile() {
	fileObj, err := os.OpenFile("./file.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	defer fileObj.Close()

	fileObj.Write([]byte("zhoulin mengbi le!\n"))
	fileObj.WriteString("周琳解释不了!")
}

func readFromByBufio() {
	fileObj, err := os.OpenFile("./file.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	defer fileObj.Close()

	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello world\n")
	wr.Flush()
}

func readFromByIoutil() {
	str := "hello hello\n"
	err := ioutil.WriteFile("./file.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed,err: ", err)
		return
	}
}

func main() {
	// writeFromFile()
	// readFromByBufio()
	readFromByIoutil()
}
