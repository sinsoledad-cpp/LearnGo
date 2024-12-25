package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/fengfengzhidao/grpc/07morebyone/stream_proto/proto"

	"google.golang.org/grpc"
)

type ClientStream struct{}

// func (ClientStream) UploadFile(stream proto.ClientStream_UploadFileServer) error {
// 	for i := 0; i < 10; i++ {
// 		response, err := stream.Recv()
// 		fmt.Println(response, err)
// 	}
// 	stream.SendAndClose(&proto.Response{Text: "完毕了"})
// 	return nil
// }

func (ClientStream) UploadFile(stream proto.ClientStream_UploadFileServer) error {

	file, err := os.OpenFile("./../static/x.png", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	var index int
	for {
		index++
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		writer.Write(response.Content)
		fmt.Printf("第%d次", index)
	}
	writer.Flush()
	stream.SendAndClose(&proto.Response{Text: "完毕了"})
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	proto.RegisterClientStreamServer(server, &ClientStream{})

	server.Serve(listen)
}
