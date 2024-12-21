package main

import (
	"fmt"
	"io"
	"os"
)

func f() {
	fileObj, err := os.OpenFile("./file.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	defer fileObj.Close()

	tmpFile, err := os.OpenFile("./file.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("create tmp file failed,err:%v\n", err)
		return
	}
	defer tmpFile.Close()

	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("read from file failed,err:%v\n", err)
		return
	}
	tmpFile.Write(ret[:n])
	var s []byte
	s = []byte{'c'}
	tmpFile.Write(s)

	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			tmpFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Printf("read from file failed,err:%v\n", err)
			return
		}
		tmpFile.Write(x[:n])
	}
	if _, err := os.Stat("./file.txt"); err == nil {
		if err := os.Remove("./file.txt"); err != nil {
			fmt.Printf("failed to remove existing file:%v\n",err)
		}
	}
	if err:=os.Rename("./file.tmp", "./file.txt");err!=nil{
		fmt.Printf("failed to rename file:%v\n",err)
	}
}
func main() {
	f()
}
