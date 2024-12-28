package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/kuanshenshuo/grpc_new/01_hello_world/hello_server/proto"
	"google.golang.org/grpc"
)

// hello server
type server struct {
	pb.UnimplementedSayHelloServer
}

// 业务
func (s *server) SayHello(cxt context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 获取元数据信息
	// md, ok := metadata.FromIncomingContext(cxt)
	// if !ok {
	// 	return nil, errors.New("获取元数据失败")
	// }
	// var appId string
	// var appKey string
	// if v, ok := md["appid"]; ok {
	// 	appId = v[0]
	// }
	// if v, ok := md["appkey"]; ok {
	// 	appKey = v[0]
	// }
	// fmt.Println("appId:", appId)
	// fmt.Println("appKey:", appKey)
	// if appId != "kuangshen" || appKey != "123123" {
	// 	return nil, errors.New("token 不正确")
	// }

	fmt.Println("收到客户端请求:", req.RequestName)
	return &pb.HelloResponse{ResponseMsg: "hello " + req.RequestName}, nil
}

func main() {
	//====:
	//TSL认证
	//====:
	//两个参数分别是 cretFile ，keyFile
	//自签名证书文件和私钥文件
	// 注意请用绝对路径
	// creds, _ := credentials.NewServerTLSFromFile("D:\\Code\\Go\\src\\github.com\\kuanshenshuo\\grpc_new\\01_hello_world\\key\\test.pem", "D:\\Code\\Go\\src\\github.com\\kuanshenshuo\\grpc_new\\01_hello_world\\key\\test.key")
	// 开启端口
	listen, _ := net.Listen("tcp", ":8080")
	// 创建grpc server
	grpcServer := grpc.NewServer()
	// grpcServer := grpc.NewServer(grpc.Creds(creds))
	// grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	// 注册服务(在grpc server中注册我们自己编写的服务)
	pb.RegisterSayHelloServer(grpcServer, &server{})
	// 开启服务
	err := grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
