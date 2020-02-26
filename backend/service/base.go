package service

import (
	"context"
	"log"
	"sync"
	"time"

	sdk "agones.dev/agones/sdks/go"
	"github.com/seanhagen/endless_stream/backend/endless"
	"github.com/seanhagen/endless_stream/backend/game"
	"github.com/seanhagen/endless_stream/backend/grpc"
	g "google.golang.org/grpc"
)

type Srv struct {
	games     map[string]*game.Game
	l         *sync.Mutex
	cancels   map[string]context.CancelFunc
	entityCol game.EntityCollection
	sdk       *sdk.SDK
}

// Setup ...
func Setup(ctx context.Context, srv *grpc.Base, ec game.EntityCollection, sdk *sdk.SDK) (*Srv, error) {
	// TODO:
	//   - return svc so that it can be gracefully shutdown
	//   - add function to svc for graceful shutdown
	//   - add method to game that gracefully shutsdown the game
	svc := &Srv{
		games:     map[string]*game.Game{},
		cancels:   map[string]context.CancelFunc{},
		l:         &sync.Mutex{},
		entityCol: ec,
		sdk:       sdk,
	}

	srv.RegisterHandler(func(s *g.Server) {
		endless.RegisterGameServer(s, svc)
	})

	go svc.background()

	return svc, nil
}

// Shutdown ...
func (s *Srv) Shutdown(ctx context.Context) {
	for _, g := range s.games {
		g.Shutdown(ctx)
	}
}

// cleanup ...
func (s *Srv) background() {
	tick := time.NewTicker(time.Second * 10)
	health := time.NewTicker(time.Second * 2)
	// on a timer, cleanup games that aren't running any more
	for {
		select {
		case <-health.C:
			if err := s.sdk.Health(); err != nil {
				log.Printf("unable to send health check: %v", err)
			}
		case t := <-tick.C:
			for id, g := range s.games {
				if !g.IsRunning(t) {
					s.l.Lock()
					ctx := context.Background()
					ctx, cancel := context.WithTimeout(ctx, time.Second*3)
					g.Shutdown(ctx)
					cancel()

					g = nil
					delete(s.games, id)

					if c, ok := s.cancels[id]; ok {
						c()
					}
					delete(s.cancels, id)

					s.l.Unlock()
				}
			}
		default:
		}
	}

}
