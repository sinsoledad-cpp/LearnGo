protoc -I ./service_proto --go_out=plugins=grpc:./service_proto ./service_proto/common.proto
protoc -I ./service_proto --go_out=plugins=grpc:./service_proto ./service_proto/order.proto
protoc -I ./service_proto --go_out=plugins=grpc:./service_proto ./service_proto/video.proto