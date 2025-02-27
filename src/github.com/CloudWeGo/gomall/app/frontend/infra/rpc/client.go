package rpc

import (
	"sync"

	"github.com/CloudWeGo/gomall/app/frontend/conf"
	frontendUtils "github.com/CloudWeGo/gomall/app/frontend/utils"
	"github.com/CloudWeGo/gomall/common/clientsuite"
	"github.com/CloudWeGo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/CloudWeGo/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/CloudWeGo/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/CloudWeGo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/CloudWeGo/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
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
