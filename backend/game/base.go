package game

import (
	"context"

	"github.com/seanhagen/endless_stream/backend/endless"
	"github.com/seanhagen/endless_stream/backend/server"
	"google.golang.org/grpc"
)

// Setup ...
func Setup(ctx context.Context, srv *server.Base) error {
	// TODO:
	//   - return svc so that it can be gracefully shutdown
	//   - add function to svc for graceful shutdown
	//   - add method to game that gracefully shutsdown the game
	svc := &Srv{
		games:   map[string]*Game{},
		cancels: map[string]context.CancelFunc{},
	}

	srv.RegisterHandler(func(s *grpc.Server) {
		endless.RegisterGameServer(s, svc)
	})

	go svc.cleanup()

	return nil
}
