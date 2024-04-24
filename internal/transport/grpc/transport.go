package grpc

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/seanhagen/endless_stream/internal/observability"
	"google.golang.org/grpc"
)

// Transport ...
type Transport struct {
	conf      internalConfig
	log       observability.Logger
	running   bool
	stopCtxFn context.CancelFunc
}

// BuildTransport ...
func BuildTransport(_ context.Context, conf Config) (*Transport, error) {
	ic, err := conf.toInternal()
	if err != nil {
		return nil, err
	}
	return &Transport{conf: ic, log: ic.logger}, nil
}

// Start ...
func (tspt *Transport) Start(ctx context.Context) error {
	options := tspt.conf.options()
	srv := grpc.NewServer(options...)

	for _, svc := range tspt.conf.services {
		svc.Register(srv)
	}

	tspt.running = true

	ctx, cancelFn := context.WithCancel(ctx)
	tspt.stopCtxFn = cancelFn

	go tspt.launchServer(ctx, srv)

	return nil
}

// launchServer ...
func (tspt *Transport) launchServer(ctx context.Context, srv *grpc.Server) {
	errChan := make(chan error)

	go tspt.serve(srv, errChan)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)

	defer func() {
		tspt.running = false
	}()

	select {
	case err := <-errChan:
		tspt.errorStop(ctx, err)

	case sig := <-sigChan:
		tspt.sigStop(ctx, srv, sig)

	case <-ctx.Done():
		tspt.ctxStop(ctx, srv)
	}
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
	srv.GracefulStop()
	tspt.log.Info(
		ctx,
		"graceful shutdown complete",
		observability.ToAttr("shutdown", observability.DurationValue(time.Since(now))),
	)
	// fmt.Printf("graceful shutdown complete in %s\n", time.Since(now))
}

// serve ...
func (tspt *Transport) serve(srv *grpc.Server, errChan chan error) {
	err := srv.Serve(tspt.conf.listener)
	if err != nil {
		errChan <- err
	}
}

// Running ...
func (tspt *Transport) Running() bool {
	return tspt.running
}

// Stop ...
func (tspt *Transport) Stop() error {
	tspt.stopCtxFn()
	return nil
}
