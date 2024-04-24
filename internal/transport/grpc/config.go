package grpc

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/seanhagen/endless_stream/internal/observability"
	"github.com/seanhagen/endless_stream/internal/observability/logs"
	"google.golang.org/grpc"
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
	StreamInterceptors []StreamInterceptorFn

	// UnaryInterceptors ...
	UnaryInterceptors []UnaryInterceptorFn

	// DialOpts ...
	DialOpts []grpc.DialOption

	// StatsHandler *stats.Handler

	// KeepAlive KeepAliveConfig

	// Max Maximums

	// EnabledSharedWriteBuffer bool

	// UnknownServiceHandler grpc.StreamHandler
}

type internalConfig struct {
	listener      net.Listener
	logger        observability.Logger
	serverOptions serverOpts
	services      []Service
}

// toInternal ...
func (conf Config) toInternal() (internalConfig, error) {
	ic := internalConfig{
		serverOptions: serverOpts{},
		services:      conf.Services,
		logger:        conf.Logger,
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

	if ic.logger == nil {
		ic.logger = logs.New(&logs.Config{Out: io.Discard})
	}

	return ic, nil
}

// options ...
func (iConf internalConfig) options() []grpc.ServerOption {
	opts := []grpc.ServerOption{}

	for _, v := range iConf.serverOptions {
		opts = append(opts, v)
	}

	return opts
}
