# 电商项目

## 使用

docker compose up -d
docker compose down

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