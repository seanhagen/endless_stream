package server

import (
	"context"
	"time"
)

// Shutdown gracefully shuts down all services
func (ba *Base) Shutdown(ctx context.Context) error {
	ba.Logger.Printf("Recieved signal to shutdown, allowing 5 seconds for graceful shutdown.")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := ba.grpc.grpcDial.Close()
	if err != nil {
		ba.Logger.Printf("Error closing GRPC connection: %v", err)
	}

	ba.grpc.srv.Stop()
	ba.grpc.cancel()

	// err = ba.grpc.httpSrv.Shutdown(ctx)
	// if err != nil {
	//   ba.Logger.Printf("unable to shutdown grpc: %v", err)
	// }

	// if ba.erRep != nil {
	//   if err = ba.erRep.Close(); err != nil {
	//     ba.Logger.Printf("unable to close error reporter: %v", err)
	//   }
	// }

	// if ba.stTr != nil {
	//   ba.stTr.Close()
	// }

	return nil
}
