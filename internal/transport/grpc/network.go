package grpc

import "net"

const (
	// DefaultGRPCPort ...
	DefaultGRPCPort int32 = 8000
	// DefaultHTTPPort ...
	DefaultHTTPPort int32 = 8080
)

// NetworkConfig ...
type NetworkConfig interface {
	apply(*internalConfig) error
}

// WithGrpcOnly ...
func WithGrpcOnly(port int32) NetworkConfig {
	return nil
}

// WithSharedGrpcGatewayPort ...
func WithSharedGrpcGatewayPort(port int32) NetworkConfig {
	return nil
}

// WithSeparateGrpcGatewayPort ...
func WithSeparateGrpcGatewayPort(grpcPort, httpPort int32) NetworkConfig {
	return nil
}

// WithCustomListener ...
func WithCustomListener(listen net.Listener) NetworkConfig {
	return customListener{listen}
}

type customListener struct {
	listen net.Listener
}

// apply ...
func (cl customListener) apply(conf *internalConfig) error {
	conf.listener = cl.listen
	return nil
}
