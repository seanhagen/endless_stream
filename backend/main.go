package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	sdk "agones.dev/agones/sdks/go"
	"github.com/gobuffalo/packr/v2"
	"github.com/seanhagen/endless_stream/backend/game"
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

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format("2006-01-02T15:04:05.999") + " [DEBUG] " + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
	log.Println("Starting server")

	// sdk, err := agones.NewSDK()
	// if err != nil {
	// 	log.Fatalf("failed to get agones sdk: %v", err)
	// }
	var sdk *sdk.SDK

	ctx := context.Background()
	ctx, cnl := context.WithCancel(ctx)
	defer cnl()

	srv, err := setup(ctx)
	if err != nil {
		log.Fatalf("Unable to setup server: %v", err)
	}

	scr, err := scripts(ctx)
	if err != nil {
		log.Fatalf("Unable to load game scripts: %v", err)
	}

	ent, err := entities(ctx)
	if err != nil {
		log.Fatalf("Unable to load game entities: %v", err)
	}

	ec, err := game.SetupEntityCollection(scr, ent)
	if err != nil {
		log.Fatalf("Unable to setup all game entities: %v", err)
	}

	svc, err := service.Setup(ctx, srv, ec, sdk)
	if err != nil {
		log.Fatalf("Unable to initialize game server: %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)

	go func() {
		// t, err := sdk.GameServer()
		// if err != nil {
		// 	log.Printf("Unable to get game server info from Agones")
		// 	sigChan <- syscall.SIGTERM
		// 	return
		// }

		if sdk != nil {
			log.Print("Marking this server as ready")
			if err := sdk.Ready(); err != nil {
				log.Printf("Could not send ready message")
				sigChan <- syscall.SIGTERM
				return
			}
		}

		err = srv.Start(ctx, cnl)
		if err != nil {
			log.Printf("Error starting or shutting down app: %v", err)
			sigChan <- syscall.SIGTERM
		}
	}()

	_ = <-sigChan
	if sdk != nil {
		if err = sdk.Shutdown(); err != nil {
			log.Printf("Unable to send shutdown message: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*15)
	svc.Shutdown(ctx)
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("")
	}
	cancel()

	log.Printf("server shutdown complete")
}

func scripts(ctx context.Context) (*packr.Box, error) {
	b := packr.New("scripts", "./scripts")
	return b, nil
}

func entities(ctx context.Context) (*packr.Box, error) {
	b := packr.New("entities", "./entities")
	return b, nil
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
