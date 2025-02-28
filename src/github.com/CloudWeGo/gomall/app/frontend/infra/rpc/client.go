package rpc

import (
	"context"
	"sync"

	"github.com/CloudWeGo/gomall/app/frontend/conf"
	frontendUtils "github.com/CloudWeGo/gomall/app/frontend/utils"
	"github.com/CloudWeGo/gomall/common/clientsuite"
	"github.com/CloudWeGo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/CloudWeGo/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/CloudWeGo/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/CloudWeGo/gomall/rpc_gen/kitex_gen/product"
	"github.com/CloudWeGo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/CloudWeGo/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
	ServiceName    = frontendUtils.ServiceName
	RegistryAddr   = conf.GetConf().Hertz.RegistryAddr
	err            error
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
	})
}

func initUserClient() {
	// r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	// frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user",
		// client.WithResolver(r),
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	)
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	//熔断策略
	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return circuitbreak.RPCInfo2Key(ri)
	})

	cbs.UpdateServiceCBConfig(
		"shop-frontend/product/GetProduct",
		circuitbreak.CBConfig{
			Enable:    true,
			ErrRate:   0.5,
			MinSample: 2,
		})

	// var opts []client.Option
	// r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	// frontendUtils.MustHandleError(err)
	// opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product",
		// opts...,
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
		//熔断策略
		client.WithCircuitBreaker(cbs),
		// Fallback
		client.WithFallback(
			fallback.NewFallbackPolicy(
				fallback.UnwrapHelper(
					func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
						methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
						if err == nil {
							return resp, err
						}
						if methodName != "ListProducts" {
							return resp, err
						}
						return &product.ListProductsResp{
							Products: []*product.Product{
								{
									Price:       6.6,
									Id:          3,
									Picture:     "/static/image/t-shirt.jpeg",
									Name:        "T-Shirt",
									Description: "CloudWeGo T-Shirt",
								},
							},
						}, nil
					}))),
	)
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	// var opts []client.Option
	// r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	// frontendUtils.MustHandleError(err)
	// opts = append(opts, client.WithResolver(r))
	CartClient, err = cartservice.NewClient("cart",
		// opts...,
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	)
	frontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	// var opts []client.Option
	// r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	// frontendUtils.MustHandleError(err)
	// opts = append(opts, client.WithResolver(r))
	CheckoutClient, err = checkoutservice.NewClient("checkout",
		// opts...,
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	)
	frontendUtils.MustHandleError(err)
}

func initOrderClient() {
	// var opts []client.Option
	// r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	// frontendUtils.MustHandleError(err)
	// opts = append(opts, client.WithResolver(r))
	OrderClient, err = orderservice.NewClient("order",
		// opts...,
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	)
	frontendUtils.MustHandleError(err)
}
