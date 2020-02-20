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

	log.Printf("Starting server (listening on: %v %v)", ba.listen.Addr().Network(), ba.listen.Addr().String())

	sigChan := make(chan os.Signal, 1)
	errChan := make(chan error)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)

	ba.setupHTTP()

	// // tcp muxer
	// tcpMux := cmux.New(ba.listen)

	// // dispatcher
	// grpcL := tcpMux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	// httpL := tcpMux.Match(cmux.HTTP1Fast())
	// httpA := tcpMux.Match(cmux.Any()) //

	// go func() {
	// 	srv := &http.Server{
	// 		Handler: fallback{},
	// 	}
	// 	srv.Serve(httpA)
	// }()

	// go func() {
	// 	// log.Printf("starting server https://%v", httpL.Addr().String())
	// 	// ba.httpSrv.Addr = httpL.Addr().String()

	// 	// err := ba.httpSrv.Serve(httpL)

	// 	// err := ba.httpSrv.ServeTLS(httpL, "/certs/cert.pem", "/certs/key.pem")
	// 	// log.Printf("starting server https://%v", ba.listen.Addr().String())
	// 	// err := ba.httpSrv.ServeTLS(ba.listen, "/certs/cert.pem", "/certs/key.pem")

	// 	err := ba.httpSrv.Serve(ba.listen)

	// 	if err != nil {
	// 		errChan <- err
	// 	}
	// 	log.Printf("http server returned")
	// }()

	go func() {
		// log.Printf("starting server grpc://%v", grpcL.Addr().String())
		// err := ba.srv.Serve(grpcL)
		log.Printf("starting server grpc://%v", ba.listen.Addr().String())
		err := ba.srv.Serve(ba.listen)
		if err != nil {
			errChan <- err
		}
		log.Printf("grpc server returned")
	}()

	// go ba.grpc.Start(ctx, errChan)

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

// // Start ...
// func (s *grpcServer) Start(ctx context.Context, errChan chan<- error) {
// 	ctx, cancel := context.WithCancel(ctx)
// 	s.cancel = cancel
// 	if s.srv == nil {
// 		errChan <- fmt.Errorf("grpc server must be initialized before calling start")
// 	}

// 	// tcp muxer
// 	// tcpMux := cmux.New(s.listen)

// 	// // dispatcher
// 	// grpcL := tcpMux.MatchWithWriters(
// 	// 	cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
// 	// httpL := tcpMux.Match(cmux.HTTP1Fast())

// 	go func() {
// 		log.Printf("starting GRPC server listening on %v", s.listen.Addr().String())
// 		if err := s.srv.Serve(s.listen); err != nil {
// 			errChan <- err
// 		}
// 	}()

// 	// go func() {
// 	// 	log.Printf("starting HTTP server listening on %v", httpL)
// 	// 	if err := s.httpSrv.Serve(httpL); err != nil {
// 	// 		errChan <- err
// 	// 	}
// 	// }()

// 	// done := make(chan bool)
// 	// go func() {
// 	// 	err := tcpMux.Serve()
// 	// 	if err != nil {
// 	// 		errChan <- err
// 	// 	}
// 	// 	done <- true
// 	// }()

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			errChan <- ctx.Err()
// 			return
// 			// case <-done:
// 			// 	return
// 		}
// 	}
// }
