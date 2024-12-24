package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"github.com/fengfengzhidao/grpc/04duo/grpc_proto/duo_proto"
	"log"
	"net"
)

type VideoServer struct {
}

func (VideoServer) Look(ctx context.Context, request *duo_proto.Request) (res *duo_proto.Response, err error) {
	fmt.Println("video:", request)
	return &duo_proto.Response{
		Name: "fengfeng",
	}, nil
}

type OrderServer struct {
}

func (OrderServer) Buy(ctx context.Context, request *duo_proto.Request) (res *duo_proto.Response, err error) {
	fmt.Println("order:", request)
	return &duo_proto.Response{
		Name: "fengfeng",
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	duo_proto.RegisterVideoServiceServer(s, &VideoServer{})
	duo_proto.RegisterOrderServiceServer(s, &OrderServer{})
	fmt.Println("grpc server程序运行在：8080")
	err = s.Serve(listen)
}
