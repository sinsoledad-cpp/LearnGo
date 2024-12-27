package main

import (
	"fmt"
	"log"
	"net"

	"github.com/fengfengzhidao/grpc/08morebymore/stream_proto/proto"

	"google.golang.org/grpc"
)

type BothStream struct{}

func (BothStream) Chat(stream proto.BothStream_ChatServer) error {
	for i := 0; i < 10; i++ {
		request, _ := stream.Recv()
		fmt.Println(request)
		stream.Send(&proto.Response{
			Text: "你好",
		})
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	proto.RegisterBothStreamServer(server, &BothStream{})

	server.Serve(listen)
}
