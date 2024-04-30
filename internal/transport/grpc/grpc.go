//go:generate go-enum --marshal --lower --names --values --ptr

// Package grpc contains wrappers & helpers for setting up GRPC services.
package grpc

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// DefaultTimeout ...
const DefaultTimeout time.Duration = time.Second * 30

// Service ...
type Service interface {
	// Register is because service needs to be able to call the correct
	// proto register method; ie proto.RegisterHexServer
	Register(grpc.ServiceRegistrar)

	// RegisterGateway is for services that want to register a
	// grpc-gateway to handle requests. Services can use
	// RegisterXHandlerFromEndpoint or RegisterXHandler.
	RegisterGateway(context.Context, *runtime.ServeMux)
}

// ENUM(Network,TLS,StatsHandler).
type serverOptionKey uint32

type serverOpts map[serverOptionKey]grpc.ServerOption

// StreamInterceptorFn ...
type StreamInterceptorFn grpc.StreamServerInterceptor

// UnaryInterceptorFn ...
type UnaryInterceptorFn grpc.UnaryServerInterceptor

type (
	// EnforcementPolicy ...
	EnforcementPolicy interface{}
	// KeepAliveParams ...
	KeepAliveParams interface{}
)

// KeepAliveConfig ...
type KeepAliveConfig struct {
	// Enforcement ...
	Enforcement EnforcementPolicy
	// Params ...
	Params KeepAliveParams
}

// Maximums ...
type Maximums struct {
	// ConcurrentStreams ...
	ConcurrentStreams uint32
	// MaxHeaderListSize ...
	MaxHeaderListSize uint32
	// MaxRecvMsgSize ...
	MaxRecvMsgSize uint32
	// MaxSendMsgSize ...
	MaxSendMsgSize uint32
	// ReadBufferSize ...
	ReadBufferSize uint32
	// WriteBufferSize ...
	WriteBufferSize uint32
}
