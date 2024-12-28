package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/kuanshengshuo/grpc_new/01_hello_world/hello_server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连接到server端,此处禁用安全传输,没有加密和验证
	// conn, err := grpc.Dial(":8080", grpc.WithInsecure())//老方法
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials())) //新方法
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 建立连接
	client := pb.NewSayHelloClient(conn)
	// 执行rpc调用(这个方法在服务器端来实现并返回结果)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(resp.GetResponseMsg())
}
