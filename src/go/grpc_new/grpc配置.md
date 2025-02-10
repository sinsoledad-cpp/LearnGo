# 配置grpc

## 下载protocol buffers

[protobuf-github](https://github.com/protocolbuffers/protobuf/releases)

Protocol bufers，通常称为 Protobuf，是 Google 开发的一种协议，用于允许对结构化数据进行序列化和反序列化。它在开发程序以通过网络相互通信或存储数据时很有用。谷歌开发它的目的是提供一种比XML更好的方式来通信。
我们将找到所有操作系统的所有 zip 文件。基于您的操作系统位版本(64位或 32位)。您必须下载特定的。
配置环境变量:D:\Environment\protoc-21.9-win64\bin
最后，我们将检查它是否有效。打开命令提示符，输入“protoc“命令

## 安装gRPC的核心库

`go get google.golang.org/grpc`

## go代码生成工具

上面安装的是protocol编译器。它可以生成各种不同语言的代码。因此，除了这个编译器，我们还需要配合各个语言的代码生成工具。对于Golang来说，称为protoc-gen-g0。不过在这儿有个小小的坑，github.com/golang/protobuf/protoc-gen-go 和 goog1e.golang.org/protobuf/cmd/protoc-gen-g0 是不同的。区别在于前者是旧版本，后者是go0gle接管后的新版本，他们之间的API是不同的，也就是说用于生成的命令，以及生成的文件都是不一样的。因为目前的QRPC-q0源码中的example用的是后者的生成方式，为了与时俱进，我们也采取最新的方式。你需要安装两个库:

`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`


因为这些文件在安装grpc的时候，已经下载下来了，因此使用insta11命令就可以了，而不需要使用get 命令。然后你看你的$GOWORKS/bin路径，应该有标1和2的两个文件:

**protoc-gen-go.exe**

**protoc-gen-go-grpc.exe**

官方网站：https://grpc.io/

底层协议：

* HTTP2: https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md
* GRPC-WEB ： https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-WEB.md