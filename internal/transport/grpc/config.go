package grpc

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/seanhagen/endless_stream/internal/observability"
	"github.com/seanhagen/endless_stream/internal/observability/logs"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/stats"
)

type Config struct {
	// Logger ...
	Logger observability.Logger

	// Network controls how the GRPC server listens for traffic. There
	// are four options:
	//
	//  - GRPC only
	//  - GRPC with grpc-gateway, sharing the same port
	//  - GRPC with grpc-gateway, with different ports
	//  - a custom listener
	//
	// If no option is provided, the default is GRPC only on the port
	// specified by DefaultGRPCPort.
	Network NetworkConfig

	// TLS ...
	TLS TLSConfig

	// Timeout ...
	Timeout time.Duration

	// Services ...
	Services []Service

	// StreamInterceptors ...
	StreamInterceptors []grpc.StreamServerInterceptor

	// UnaryInterceptors ...
	UnaryInterceptors []grpc.UnaryServerInterceptor

	// DialOpts ...
	DialOpts []grpc.DialOption

	StatsHandler stats.Handler

	// KeepAlive KeepAliveConfig

	// Max Maximums

	// EnabledSharedWriteBuffer bool

	// UnknownServiceHandler grpc.StreamHandler
}

type internalConfig struct {
	tcpMux          cmux.CMux
	listener        net.Listener
	gatewayListener net.Listener
	useGateway      bool
	separatePorts   bool

	shutdown func(context.Context)

	logger        observability.Logger
	serverOptions serverOpts
	services      []Service

	unary  []grpc.UnaryServerInterceptor
	stream []grpc.StreamServerInterceptor
}

// toInternal ...
func (conf Config) toInternal() (internalConfig, error) {
	ic := internalConfig{
		serverOptions: serverOpts{},
		services:      conf.Services,
		logger:        conf.Logger,
		unary:         conf.UnaryInterceptors,
		stream:        conf.StreamInterceptors,
	}

	var err error
	if conf.Network != nil {
		err = conf.Network.apply(&ic)
	} else {
		err = WithGrpcOnly(DefaultGRPCPort).apply(&ic)
	}
	if err != nil {
		return ic, fmt.Errorf("unable to configure network: %w", err)
	}

	if conf.TLS != nil {
		err = conf.TLS.apply(&ic)
	} else {
		err = WithInsecureTLS().apply(&ic)
	}
	if err != nil {
		return ic, fmt.Errorf("unable to apply TLS configuration: %w", err)
	}

	if ic.logger == nil {
		ic.logger = logs.New(&logs.Config{Out: io.Discard})
	}

	return ic, nil
}

// options ...
func (iConf internalConfig) options() []grpc.ServerOption {
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			iConf.unary...,
		),
		grpc.ChainStreamInterceptor(
			iConf.stream...,
		),
	}

	for _, v := range iConf.serverOptions {
		opts = append(opts, v)
	}

	return opts
}
