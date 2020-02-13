package main

import (
	"context"
	"log"

	"github.com/seanhagen/endless_stream/backend/grpc"
	"github.com/seanhagen/endless_stream/backend/service"
	g "google.golang.org/grpc"
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

	err = service.Setup(ctx, srv)
	if err != nil {
		log.Fatalf("Unable to initialize game server: %v", err)
	}

	log.Printf("server setup, starting")
	err = srv.Start(ctx, cancel)
	if err != nil {
		log.Fatalf("Error starting or shutting down app: %v", err)
	}

	log.Printf("server shutdown complete")
}

func setup(ctx context.Context) (*grpc.Base, error) {
	conf := grpc.Config{
		Version:            Version,
		Build:              Build,
		Repo:               Repo,
		UnaryInterceptors:  []g.UnaryServerInterceptor{},
		StreamInterceptors: []g.StreamServerInterceptor{},
	}

	return grpc.New(ctx, conf)
}
