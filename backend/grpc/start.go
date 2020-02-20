package grpc

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	err := ba.createGRPCServer(ctx, conf)
	if err != nil {
		return err
	}

	sigChan := make(chan os.Signal, 1)
	errChan := make(chan error)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)

	go func() {
		log.Printf("starting server grpc://%v", ba.listen.Addr().String())
		err := ba.srv.Serve(ba.listen)
		if err != nil {
			errChan <- err
		}
		log.Printf("grpc server returned")
	}()

	select {
	case err := <-errChan:
		if err != nil {
			ba.Logger.Printf("encountered error during runtime: %v", err)
			x := ba.Shutdown(ctx)
			if x != nil {
				ba.Logger.Printf("unable to shutdown service after encountering error: %v", x)
			}
		}
		return err
	case <-sigChan:
		cancel()
		log.Printf("received signal to shutdown")
		return ba.Shutdown(ctx)
	}
}
