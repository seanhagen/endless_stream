package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"google.golang.org/grpc"
)

const grpcListen = "10001"

// DefaultTimeout is the timeout used when the grpc-gateway sends requests
// to the grpc server. The default is 0 -- no timeout.
var DefaultTimeout = 0 * time.Second

// GatewayHandler ...
type GatewayHandler func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error

// Handler ...
type Handler func(*grpc.Server)

type grpcConfig struct {
	vip *viper.Viper
	// trace              *stats.Tracer
	// error              *errors.Reporter
	UnaryInterceptors  []grpc.UnaryServerInterceptor
	StreamInterceptors []grpc.StreamServerInterceptor
}

type grpcServer struct {
	cancel   context.CancelFunc
	srv      *grpc.Server
	listen   net.Listener
	dopts    []grpc.DialOption
	httpSrv  *http.Server
	grpcDial *grpc.ClientConn
	// tr       *stats.Tracer
}

// registerHandler ...
func (s *grpcServer) registerHandler(ctx context.Context, fn Handler) {
	fn(s.srv)
}

// createGRPCServer creates a GRPC server from the config
func (ba *Base) createGRPCServer(ctx context.Context, conf grpcConfig) (*grpcServer, error) {
	listenPort := shouldEnv("PORT", grpcListen)
	// setup default grpc listener
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", listenPort))
	if err != nil {
		return nil, err
	}

	// setup default dial opts
	dopts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	// setup grpc dialer
	conn, err := grpc.Dial(listenPort, dopts...)
	if err != nil {
		return nil, err
	}

	log.Printf("Server setup to listen on %v", listenPort)

	srv, err := ba.setupGRPC(conf)
	if err != nil {
		return nil, err
	}

	s := &grpcServer{
		srv:      srv,
		listen:   l,
		dopts:    dopts,
		grpcDial: conn,
		// tr:       conf.trace,
	}

	return s, nil
}

// setupGRPC sets up the gRPC server
func (ba *Base) setupGRPC(grpcConf grpcConfig) (*grpc.Server, error) {
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

	return grpc.NewServer(srvOpts...), nil
}
