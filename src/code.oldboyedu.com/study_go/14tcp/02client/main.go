package main

import (
	"fmt"
	"net"

	proto "code.oldboyedu.com/study_go/14tcp/proto"
)

func main() {
	conn, err := net.Dial("tcp","127.0.0.1:30000")
	if err!=nil{
		fmt.Println("dial failled,err: ",err)
		return
	}
	defer conn.Close()

	for i:=0;i<20;i++{
		msg:="hello,hello. How are you?"
		// conn.Write([]byte(msg))
		b,err:=proto.Encode(msg)
		if err!=nil{
			fmt.Println("encode failed,err:",err)
			return
		}
		conn.Write(b)
	}
}

