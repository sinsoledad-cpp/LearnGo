package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/kuanshengshuo/grpc_new/01_hello_world/hello_server/proto"
	"google.golang.org/grpc"
)

// hello server
type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(cxt context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("收到客户端请求:", req.RequestName)
	return &pb.HelloResponse{ResponseMsg: "hello " + req.RequestName}, nil
}

func main() {
	// 开启端口
	listen, _ := net.Listen("tcp", ":8080")
	// 创建grpc server
	grpcServer := grpc.NewServer()
	// 注册服务(在grpc server中注册我们自己编写的服务)
	pb.RegisterSayHelloServer(grpcServer, &server{})
	// 开启服务
	err := grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
