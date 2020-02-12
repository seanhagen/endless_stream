package grpc

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/soheilhy/cmux"
)

// Start starts up the server. The cancel func is so that the base app can notify
// the service that it's about to shutdown
func (ba *Base) Start(ctx context.Context, cancel context.CancelFunc) error {
	conf := grpcConfig{
		vip: ba.Config,
		// trace:              ba.stTr,
		// error:              ba.erRep,
		UnaryInterceptors:  ba.unaryIntercept,
		StreamInterceptors: ba.streamIntercept,
	}

	grpc, err := ba.createGRPCServer(ctx, conf)
	if err != nil {
		return err
	}
	ba.grpc = grpc

	for _, fn := range ba.handlers {
		grpc.registerHandler(ctx, fn)
	}

	ba.Logger.Printf("Starting server")

	sigChan := make(chan os.Signal, 1)
	errChan := make(chan error)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)

	go ba.grpc.Start(ctx, errChan)

	select {
	case err := <-errChan:
		ba.Logger.Printf("encountered error during runtime: %v", err)
		x := ba.Shutdown(ctx)
		if x != nil {
			ba.Logger.Printf("unable to shutdown service after encountering error: %v", x)
		}
		return err
	case <-sigChan:
		cancel()
		return ba.Shutdown(ctx)
	}
}

// Start ...
func (s *grpcServer) Start(ctx context.Context, errChan chan<- error) {
	ctx, cancel := context.WithCancel(ctx)
	s.cancel = cancel

	if s.srv == nil {
		errChan <- fmt.Errorf("grpc server must be initialized before calling start")
	}

	// tcp muxer
	tcpMux := cmux.New(s.listen)

	// dispatcher
	grpcL := tcpMux.MatchWithWriters(
		cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	// httpL := tcpMux.Match(cmux.HTTP1Fast())

	go func() {
		if err := s.srv.Serve(grpcL); err != nil {
			errChan <- err
		}
	}()

	// go func() {
	//   if err := s.httpSrv.Serve(httpL); err != nil {
	//     errChan <- err
	//   }
	// }()

	done := make(chan bool)
	go func() {
		err := tcpMux.Serve()
		if err != nil {
			errChan <- err
		}
		done <- true
	}()

	for {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
			return
		case <-done:
			return
		}
	}
}
