package rpc

import (
	"sync"

	"github.com/CloudWeGo/gomall/common/clientsuite"

	"github.com/CloudWeGo/gomall/app/cart/conf"
	cartUtils "github.com/CloudWeGo/gomall/app/cart/utils"
	"github.com/CloudWeGo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
	err           error
)

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	// var opts []client.Option
	// r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	// cartUtils.MustHandleError(err)
	// opts = append(opts, client.WithResolver(r))

	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartUtils.MustHandleError(err)
}
