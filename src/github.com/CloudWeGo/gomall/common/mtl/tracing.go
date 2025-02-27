package mtl

import (
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

func InitTracing(serviceName string) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		// provider.WithExportEndpoint("localhost:4317"),//可不写默认就是这个
		provider.WithInsecure(),
		provider.WithEnableMetrics(false), //关闭性能指标功能，因为我们使用prometheus
	)
	return p
}
