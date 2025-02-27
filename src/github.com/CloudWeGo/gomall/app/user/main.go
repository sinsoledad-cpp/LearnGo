package main

import (
	"context"
	"net"
	"time"

	"github.com/CloudWeGo/gomall/common/mtl"
	"github.com/CloudWeGo/gomall/common/serversuite"

	"github.com/CloudWeGo/gomall/app/user/biz/dal"
	"github.com/CloudWeGo/gomall/app/user/conf"
	"github.com/CloudWeGo/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)

func main() {
	err := godotenv.Load()
	if err != nil {
		klog.Error(err.Error())
	}

	mtl.InitMetric(ServiceName, conf.GetConf().Kitex.MetricsPort, RegistryAddr)
	// traceing-OpenTelemetry
	p := mtl.InitTracing(ServiceName)
	defer p.Shutdown(context.Background())
	dal.Init()

	opts := kitexInit()

	svr := userservice.NewServer(new(UserServiceImpl), opts...)

	err = svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	opts = append(opts, server.WithSuite(
		serversuite.CommonServerSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		},
	))

	// // service info
	// opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
	// 	ServiceName: conf.GetConf().Kitex.Service,
	// }))

	// // consul
	// r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// opts = append(opts, server.WithRegistry(r))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
