package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/seanhagen/endless_stream/internal/observability"
	"github.com/soheilhy/cmux"
)

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
	return grpcOnly{port}
}

type grpcOnly struct {
	port int32
}

// apply ...
func (gO grpcOnly) apply(conf *internalConfig) error {
	addr := fmt.Sprintf("0.0.0.0:%d", gO.port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("unable to set up listener on port %d: %w", gO.port, err)
	}

	conf.listener = listener

	conf.shutdown = func(ctx context.Context) {
		if err := listener.Close(); err != nil {
			conf.logger.Error(ctx, "unable to close GRPC listener", observability.ErrorAttr(err))
		}
	}

	return nil
}

// WithSharedGrpcGatewayPort ...
func WithSharedGrpcGatewayPort(port int32) NetworkConfig {
	return gatewaySharedPort{port}
}

type gatewaySharedPort struct {
	port int32
}

// apply ...
func (gsp gatewaySharedPort) apply(conf *internalConfig) error {
	addr := fmt.Sprintf("0.0.0.0:%d", gsp.port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("unable to set up listener on port %d: %w", gsp.port, err)
	}

	tcpMux := cmux.New(listener)

	grpcL := tcpMux.MatchWithWriters(
		cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := tcpMux.Match(cmux.HTTP1Fast())

	conf.tcpMux = tcpMux
	conf.listener = grpcL
	conf.gatewayListener = httpL
	conf.useGateway = true

	conf.shutdown = func(ctx context.Context) {
		tcpMux.Close()
	}

	return nil
}

// WithSeparateGrpcGatewayPort ...
func WithSeparateGrpcGatewayPort(grpcPort, httpPort int32) NetworkConfig {
	return separatePorts{grpcPort, httpPort}
}

type separatePorts struct {
	grpcPort int32
	httpPort int32
}

// apply(conf *internalConfig) ...
func (sp separatePorts) apply(conf *internalConfig) error {
	addr := fmt.Sprintf("0.0.0.0:%d", sp.grpcPort)
	grpcListener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("unable to set up grpc listener on port %d: %w", sp.grpcPort, err)
	}

	addr = fmt.Sprintf("0.0.0.0:%d", sp.httpPort)
	httpListener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("unable to set up http listener on port %d: %w", sp.httpPort, err)
	}

	conf.listener = grpcListener
	conf.gatewayListener = httpListener
	conf.useGateway = true

	conf.shutdown = func(ctx context.Context) {
		if err := grpcListener.Close(); err != nil {
			conf.logger.Error(ctx, "unable to close GRPC listener", observability.ErrorAttr(err))
		}

		if err := httpListener.Close(); err != nil {
			conf.logger.Error(ctx, "unable to close HTTP listener", observability.ErrorAttr(err))
		}
	}

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
	conf.shutdown = func(ctx context.Context) {
		if err := cl.listen.Close(); err != nil {
			conf.logger.Error(
				ctx,
				"error when closing custom listener",
				observability.ErrorAttr(err),
			)
		}
	}

	return nil
}
