package service

import (
	"context"
	"sync"
	"time"

	"github.com/seanhagen/endless_stream/backend/endless"
	"github.com/seanhagen/endless_stream/backend/game"
	"github.com/seanhagen/endless_stream/backend/grpc"
	g "google.golang.org/grpc"
)

type Srv struct {
	games   map[string]*game.Game
	l       *sync.Mutex
	cancels map[string]context.CancelFunc
}

// Setup ...
func Setup(ctx context.Context, srv *grpc.Base) error {
	// TODO:
	//   - return svc so that it can be gracefully shutdown
	//   - add function to svc for graceful shutdown
	//   - add method to game that gracefully shutsdown the game
	svc := &Srv{
		games:   map[string]*game.Game{},
		cancels: map[string]context.CancelFunc{},
		l:       &sync.Mutex{},
	}

	srv.RegisterHandler(func(s *g.Server) {
		endless.RegisterGameServer(s, svc)
	})

	go svc.cleanup()

	return nil
}

// cleanup ...
func (s *Srv) cleanup() {
	tick := time.NewTicker(time.Second * 10)
	// on a timer, cleanup games that aren't running any more
	for {
		select {
		case <-tick.C:
			for id, g := range s.games {
				if !g.Running {
					s.l.Lock()

					g = nil
					delete(s.games, id)

					if c, ok := s.cancels[id]; ok {
						c()
					}
					delete(s.cancels, id)

					s.l.Unlock()
				}
			}
		}
	}
}