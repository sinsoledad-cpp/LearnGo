package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/CloudWeGo/gomall/demo/demo_proto/conf"
	"github.com/CloudWeGo/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/CloudWeGo/gomall/demo/demo_proto/kitex_gen/pbapi/echoservice"
	"github.com/CloudWeGo/gomall/demo/demo_proto/middleware"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/transmeta"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {

	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	// c, err := echoservice.NewClient("demo_proto", client.WithResolver(r),
	// 	client.WithTransportProtocol(transport.GRPC),
	// 	client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	// )
	c, err := echoservice.NewClient("demo_proto", client.WithResolver(r),
		// client.WithTransportProtocol(transport.GRPC), 不使用 gRPC,在win11上
		client.WithShortConnection(), // 使用短链接
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithMiddleware(middleware.Middleware),
	)
	if err != nil {
		panic(err)
	}

	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_proto_client")
	res, err := c.Echo(ctx, &pbapi.Request{Message: "error"})
	fmt.Println(res)
	fmt.Println(err)
	// 返回err有bug
	var bizErr *kerrors.GRPCBizStatusError
	if err != nil {
		ok := errors.As(err, &bizErr)
		if ok {
			fmt.Printf("%#v\n", bizErr)
		}
		klog.Fatal(err)
	}

	fmt.Printf("%v\n", res)
}
