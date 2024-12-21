package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromFile() {
	fileObj, err := os.Open("./01file.go")
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	defer fileObj.Close()

	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("read all")
			return
		}
		if err != nil {
			fmt.Printf("read file failed,err:%v", err)
			return
		}
		fmt.Printf("read %d byte\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

func readFrombyBufio() {
	fileObj, err := os.Open("./01file.go")
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed,err: %v", err)
			return
		}
		fmt.Print(line)
	}
}

func readFromByIoutil() {
	ret, err := ioutil.ReadFile("./01file.go")
	if err != nil {
		fmt.Printf("read file failed,err:%v\n", err)
		return
	}
	fmt.Println(string(ret))
}
func main() {
	// readFromFile()
	// readFrombyBufio()
	readFromByIoutil()
}
