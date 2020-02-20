package grpc

import (
	"context"
	"log"
	"time"
)

// Shutdown gracefully shuts down all services
func (ba *Base) Shutdown(ctx context.Context) error {
	log.Printf("Recieved signal to shutdown, allowing 5 seconds for graceful shutdown.")
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	ba.srv.Stop()

	log.Printf("finished shutting down")
	return nil
}
