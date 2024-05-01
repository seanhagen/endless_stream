package grpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"github.com/seanhagen/endless_stream/internal/observability"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

// Transport ...
type Transport struct {
	conf      internalConfig
	log       observability.Logger
	grpcMux   *runtime.ServeMux
	running   bool
	stopCtxFn context.CancelFunc
}

// BuildTransport ...
func BuildTransport(ctx context.Context, conf Config) (*Transport, error) {
	ic, err := conf.toInternal()
	if err != nil {
		return nil, err
	}

	tspt := &Transport{conf: ic, log: ic.logger}
	tspt.setup()

	return tspt, nil
}

// setup ...
func (tspt *Transport) setup() {
	gmux := runtime.NewServeMux(
		// // pulls data from request to send as part of context
		// runtime.WithMetadata(metadataAnnotate),

		// // headerMatcher sends the matching headers to grpc
		// runtime.WithIncomingHeaderMatcher(incomingHeaderMatcher),

		// // sends the headers back to the HTTP client
		// runtime.WithOutgoingHeaderMatcher(outgoingHeaderMatcher),

		// changes json serializer to include empty fields with default values and to use camelCase instead
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:     false,
					EmitDefaultValues: true,
				},
			},
		),
		// runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)

	tspt.grpcMux = gmux
}

// Start ...
func (tspt *Transport) Start(ctx context.Context) error {
	tspt.log.Debug(ctx, "starting GRPC transport")

	options := tspt.conf.options()
	srv := grpc.NewServer(options...)

	tspt.log.Debug(ctx, "registering services")
	for _, svc := range tspt.conf.services {
		attr := observability.StringAttr("type", fmt.Sprintf("%T", srv))
		tspt.log.Debug(ctx, "registering service", attr)
		svc.Register(srv)
		if tspt.conf.useGateway {
			tspt.log.Debug(ctx, "registering service with grpc gateway")
			svc.RegisterGateway(ctx, tspt.grpcMux)
		}
	}
	tspt.log.Debug(ctx, "finished registering services")

	tspt.running = true
	ctx, cancelFn := context.WithCancel(ctx)
	tspt.stopCtxFn = cancelFn

	tspt.log.Debug(ctx, "launching server")
	go tspt.launchServer(ctx, srv)

	tspt.log.Debug(ctx, "done launching server, returning from Start")
	return nil
}

// launchServer ...
func (tspt *Transport) launchServer(ctx context.Context, srv *grpc.Server) {
	tspt.log.Debug(ctx, "server launcher goroutine started")
	errChan := make(chan error)

	go tspt.serveGRPC(ctx, srv, errChan)

	if tspt.conf.useGateway {
		go tspt.serveHTTP(ctx, errChan)
		go tspt.startTCPMux(ctx, errChan)
	}

	tspt.log.Debug(ctx, "server goroutines launched, awaiting signal or error")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)

	// defer func() {
	// 	tspt.running = false
	// }()

	select {
	case err := <-errChan:
		tspt.errorStop(ctx, err)

	case sig := <-sigChan:
		tspt.sigStop(ctx, srv, sig)

	case <-ctx.Done():
		tspt.ctxStop(ctx, srv)
	}
	tspt.log.Debug(ctx, "server launcher finished")
}

// errorStop ...
func (tspt *Transport) errorStop(ctx context.Context, err error) {
	tspt.log.Error(
		ctx,
		"error from running GRPC server during shutdown",
		observability.ErrorAttr(err),
	)
	// fmt.Printf("error from running GRPC server: %v\n", err)
}

// sigStop ...
func (tspt *Transport) sigStop(ctx context.Context, srv *grpc.Server, sig os.Signal) {
	tspt.log.Info(ctx, "received signal", observability.SigAttr(sig))
	// fmt.Printf("received signal %v to quit\n", sig)
	srv.Stop()
}

// ctxStop ...
func (tspt *Transport) ctxStop(ctx context.Context, srv *grpc.Server) {
	tspt.log.Info(ctx, "gracefully stopping GRPC server")
	// fmt.Printf("gracefully stopping GRPC server\n")
	now := time.Now()

	tspt.conf.shutdown(ctx)
	srv.GracefulStop()
	tspt.running = false

	tspt.log.Info(
		ctx,
		"graceful shutdown complete",
		observability.ToAttr("shutdown", observability.DurationValue(time.Since(now))),
	)
}

// startTCPMux ...
func (tspt *Transport) startTCPMux(ctx context.Context, errChan chan error) {
	if tspt.conf.separatePorts {
		return
	}

	if tspt.conf.tcpMux == nil {
		return
	}

	tspt.log.Debug(ctx, "starting tcp mux listener")

	if err := tspt.conf.tcpMux.Serve(); err != nil && !errors.Is(err, net.ErrClosed) {
		tspt.log.Error(ctx, "tcp mux encountered error", observability.ErrorAttr(err))
		errChan <- err
	}
}

// serveGRPC ...
func (tspt *Transport) serveGRPC(ctx context.Context, srv *grpc.Server, errChan chan error) {
	tspt.log.Debug(ctx, "launching GRPC listener")
	err := srv.Serve(tspt.conf.listener)
	if err != nil {
		tspt.log.Error(ctx, "grpc server encountered error", observability.ErrorAttr(err))
		errChan <- err
	}
}

// serveHTTP ...
func (tspt *Transport) serveHTTP(ctx context.Context, errChan chan error) {
	tspt.log.Debug(ctx, "launching GRPC gateway HTTP listener")

	runtime.DefaultContextTimeout = DefaultTimeout

	mux := http.NewServeMux()
	mux.Handle("/", tspt.grpcMux)

	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))

	corsConf := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedHeaders: []string{
			"DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since",
			"Cache-Control", "Content-Type", "Content-Range", "Range", "Authorization",
			"X-Host", "X-HTTP-Host", "X-Request-ID", "X-Server-Name", "X-Request-URI",
			"X-User-Agent", "X-Referrer",
		},
		ExposedHeaders: []string{
			"DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since",
			"Cache-Control", "Content-Type", "Content-Range", "Range", "Authorization",
		},
	})

	srv := &http.Server{
		Handler:      corsConf.Handler(mux),
		ReadTimeout:  DefaultTimeout,
		WriteTimeout: DefaultTimeout,
		IdleTimeout:  DefaultTimeout,
	}

	err := srv.Serve(tspt.conf.gatewayListener)
	if err != nil && tspt.running {
		tspt.log.Error(ctx, "http server encountered error", observability.ErrorAttr(err))
		errChan <- err
	}
}

// Running ...
func (tspt *Transport) Running() bool {
	return tspt.running
}

// Stop ...
func (tspt *Transport) Stop() error {
	tspt.log.Debug(context.Background(), "transport asked to stop")
	tspt.stopCtxFn()
	return nil
}
