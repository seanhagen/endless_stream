package grpc

import (
	"context"
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"google.golang.org/grpc"
)

const grpcListen = "8000"

// createGRPCServer creates a GRPC server from the config
func (ba *Base) createGRPCServer(ctx context.Context, conf grpcConfig) error {
	listenPort := shouldEnv("PORT", grpcListen)
	// setup default grpc listener
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", listenPort))
	if err != nil {
		return err
	}
	ba.listen = l

	err = ba.setupGRPC(conf)
	if err != nil {
		return err
	}

	return nil
}

// setupGRPC sets up the gRPC server
func (ba *Base) setupGRPC(grpcConf grpcConfig) error {
	logrusEntry := logrus.NewEntry(ba.Logger)

	logOpts := []grpc_logrus.Option{}
	// Make sure that log statements internal to grpc library are logged using the logrus Logger as well.
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	unaryInterceptors := []grpc.UnaryServerInterceptor{
		// middleware.MetadataUnary(grpcConf.vip),
		grpc_logrus.UnaryServerInterceptor(logrusEntry, logOpts...),
		// middleware.TraceUnary(),
		// middleware.ErrorUnary(grpcConf.vip),
		// middleware.ValidationUnary(),
		// middleware.InsertHeaderFilter(),
		// middleware.UnaryPanicHandler,
	}
	unaryInterceptors = append(unaryInterceptors, grpcConf.UnaryInterceptors...)

	streamInterceptors := []grpc.StreamServerInterceptor{
		// middleware.MetadataStream(grpcConf.vip),
		grpc_logrus.StreamServerInterceptor(logrusEntry, logOpts...),
		// middleware.TraceStream(),
		// middleware.ErrorStream(grpcConf.vip),
		// middleware.ValidationStream(),
		// middleware.StreamPanicHandler,
	}
	streamInterceptors = append(streamInterceptors, grpcConf.StreamInterceptors...)

	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		ba.Logger.Printf("Unable to register grpc server metric views: %v", err)
	}

	srvOpts := []grpc.ServerOption{
		grpc.StatsHandler(&ocgrpc.ServerHandler{}),
		grpc_middleware.WithUnaryServerChain(unaryInterceptors...),
		grpc_middleware.WithStreamServerChain(streamInterceptors...),
	}

	gs := grpc.NewServer(srvOpts...)
	for _, fn := range ba.handlers {
		fn(gs)
	}
	ba.srv = gs

	return nil
}
