package main

import (
	"context"
	"log"

	"github.com/seanhagen/endless_stream/backend/game"
	"github.com/seanhagen/endless_stream/backend/server"
	"google.golang.org/grpc"
)

var (
	// Version is set by the build process, contains semantic version
	Version string
	// Build is set by the build process, contains sha tag of build
	Build string
	// Repo is set by the build process, contains the repo where the code for this binary was built from
	Repo string
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	srv, err := setup(ctx)
	if err != nil {
		log.Fatalf("Unable to setup server: %v", err)
	}

	err = game.Setup(ctx, srv)
	if err != nil {
		log.Fatalf("Unable to initialize game server: %v", err)
	}

	err = srv.Start(ctx, cancel)
	if err != nil {
		log.Fatalf("Error starting or shutting down app: %v", err)
	}

	log.Printf("server shutdown complete")
}

func setup(ctx context.Context) (*server.Base, error) {
	conf := server.Config{
		Version:            Version,
		Build:              Build,
		Repo:               Repo,
		UnaryInterceptors:  []grpc.UnaryServerInterceptor{},
		StreamInterceptors: []grpc.StreamServerInterceptor{},
	}

	return server.New(ctx, conf)
}
