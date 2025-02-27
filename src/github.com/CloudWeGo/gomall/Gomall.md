# 电商项目

## 使用

docker compose up -d
docker compose down
go run .

## 开发环境

Visual Studio Code
Plugins
- Go
- Golang Tools
- Docker
- MySQL
- Material lcon Theme
- YAML
- vscode-proto3
- Makefile Tools

Oh My Zsh
Plugins
- zsh-syntax-highlighting
- zsh-autosuggestions

### Hertz

**导入:**`go get -u github.com/cloudwego/hertz`

[官网](https://www.cloudwego.io/zh/docs/hertz/getting-started/)

## 手脚架

### IDL

https://en.wikipedia.org/wiki/IDL (programming_language)

### Thrift

[Cwgo-github](https://github.com/cloudwego/cwgo)

[thriftgo-github](https://github.com/cloudwego/thriftgo)

[idl-官网](https://thrift.apache.org/docs/idl)

[idl-第三方文档](https://diwakergupta.github.io/thrift-missing-guide/)

`go install github.com/cloudwego/thriftgo@latest`

[Gitee 极速下载/Thriftgo](https://gitee.com/mirrors/Thriftgo)

`go install github.com/cloudwego/cwgo@latest`

[自动补全](https://www.cloudwego.io/zh/docs/cwgo/tutorials/auto-completion/)

[Cwgo](https://www.cloudwego.io/docs/cwgo/)

[Cwgo](https://github.com/cloudwego)

```zsh
mkdir autocomplete # You can choose any location you like
cwgo completion zsh > ./autocomplete/zsh_autocomplete
source ./autocomplete/zsh_autocomplete
```

`cwgo server --type RPC --module projects/gomall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift`

### Protobuf

[Language Guide (proto 3)](https://protobuf.dev/programming-guides/proto3/)

`wget https://github.com/protocolbuffers/protobuf/releases/download/v28.3/protoc-28.3-linux-x86_64.zip`

`unzip protoc-28.3-linux-x86_64.zip`

`sudo cp /protoc /usr/local/bin`

`sudo cp -a google /usr/local/include`

[Protocol](https://github.com/protocolbuffers/protobuf/releases/tag/v28.3)

[ProtoBuf](https://blog.csdn.net/weixin_74531333/article/details/140469169)

`cwgo server -I ../../idl --type RPC --module projects/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto`

### cloudwego

https://www.cloudwego.io/docs/hertz/tutorials/toolkit/annotation/

[hertz](https://www.cloudwego.io/docs/hertz/)

## 服务注册与服务发现 

### consul

https://hub.docker.com/r/hashicorp/consul?uuid=40228A54-EB5B-4483-B081-6E4539FF95E5

## 配置

### 方法

Env Config

Linux env
export APP_ENV=online

.env file
APP_ENV=online

Docker env
ENV GO_ENV=online

K8s env
direct
from status
from configmap
from secret

[加载.env文件的库](https://github.com/joho/godotenv)

## ORM

[gorm中文文档](https://gorm.golang.ac.cn/docs/models.html)

## 编码规范

### code style

#### style guide

https://github.com/uber-go/guide/blob/master/style.md

https://protobuf.dev/programming-guides/style/

#### tools

https://github.com/mvdan/gofumpt

https://golangci-lint.run/

### http codes

#### Http response status codes

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status

- Informational responses(100-199)
- Successful responses(200-299)
- Redirection messages(300-399)
- Client error responses (400-499)
- Server error responses(500-599)

## 特殊第三方库

模拟支付：`go get github.com/durango/go-credit-card`

![结算支付](image/结算支付.png)

[nats golang sdk](https://github.com/nats-io/nats.go)

### prometheus

[kitex prometheus](https://github.com/kitex-contrib/monitor-prometheus)

[prometheus golang sdk](https://github.com/prometheus/client_golang)

### opentelemetry

[hertz-contrib opentelemetry](https://github.com/hertz-contrib/obs-opentelemetry)

[ketix-contrib opentelemetry](https://github.com/kitex-contrib/obs-opentelemetry)

[gorm opentelemetry](https://github.com/go-gorm/opentelemetry)