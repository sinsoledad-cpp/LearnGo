package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/fengfengzhidao/grpc/06onebymore/stream_proto/proto"

	"google.golang.org/grpc"
)

type ServiceStream struct{} // 一问多答

func (ServiceStream) Fun(request *proto.Request, stream proto.ServiceStream_FunServer) error {
	fmt.Println(request)
	for i := 0; i < 10; i++ {
		stream.Send(&proto.Response{
			Text: fmt.Sprintf("第%d轮数据", i),
		})
	}
	return nil
}

func (ServiceStream) DownLoadFile(request *proto.Request, stream proto.ServiceStream_DownLoadFileServer) error {
	fmt.Println(request)
	file, err := os.Open("./../static/浮图秀图片_pic.netbian.com_20230919234843.jpg")
	if err != nil {
		return err
	}
	defer file.Close()

	for {
		buf := make([]byte, 2048)
		_, err = file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		stream.Send(&proto.FileResponse{
			Content: buf,
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
	proto.RegisterServiceStreamServer(server, &ServiceStream{})

	server.Serve(listen)
}
