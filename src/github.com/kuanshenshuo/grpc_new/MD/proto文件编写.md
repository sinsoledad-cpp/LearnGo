# proto文件编写

## 编写

```proto
// 这是在说明我们使用的是proto3语法
syntax = "proto3";

// 这部分的内容是关于最后生成的go文件是处在哪个日录哪个包中,.代表在当前目录生成，service代表了生成的go文件的包名是service.
option go_package = ".;service";

// 然后我们需要定义一个服务，在这个服务中需要有一个方法，这个方法可以接受客户端的参数，再返回服务端的响应
// 其实很容易可以看出，我们定义了一个service，称为SayHello，这个服务中有一个rpc方法，名为SayHello.
// 这个方法会发送一个HelloRequest，然后返回个HelloResponse.
service SayHello {
  rpc SayHello(HelloRequest) returns (HelloResponse) {}
}

// message关键字，其实你可以理解为Golang中的结构体。
// 这里比较特别的是变量后面的"赋值"。注意，这里并不是赋值，而是在定义这个变量在这个message中的位置。
message HelloRequest {
  string requestName = 1;
  //   int64 age = 2;
}

message HelloResponse {
  // 注意后边的数字
  string responseMsg = 1;
}
```

## 生成代码

在编写完上面的内容后，在helloworld/proto
目录下执行如下命令

protoc --go_out=. hello.proto

protoc --go-grpc_out=. hello.proto

## 服务端编写

创建 gRPC Server 对象，你可以理解为它是 Server 端的抽象对象
将 server(其包含需要被调用的服务端接口)注册到gRPCServer 的内部注册中心。这样可以在接受到请求时，通过内部的服务发现，发现该服务端接口并转接进行逻辑处理
创建 Listen，监听 TCP 端口
gRPC Server开始 lis.Accept，直到 Stop

## 客户端编写

创建与给定目标(服务端)的连接交互
创建 server的客户端对象
发送 RPC 请求，等待同步响应，得到回调后返回响应结果
输出响应结果
