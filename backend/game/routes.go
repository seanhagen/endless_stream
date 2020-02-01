package game

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/seanhagen/endless_stream/backend/endless"
)

type Srv struct {
	games   map[string]*Game
	l       *sync.Mutex
	cancels map[string]context.CancelFunc
}

// cleanup ...
func (s *Srv) cleanup() {
	tick := time.NewTicker(time.Second * 10)
	// on a timer, cleanup games that aren't running any more
	for {
		select {
		case <-tick.C:
			for id, g := range s.games {
				if !g.running {
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

// Create ...
func (s *Srv) Create(ctx context.Context, in *endless.CreateGame) (*endless.GameCreated, error) {
	var id string

	for {
		id = getGameId()
		if _, ok := s.games[id]; ok {
			continue
		}
		break
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	g, err := createGame(ctx, id)
	if err != nil {
		cancel()
		return nil, err
	}
	go func() {
		defer func() {
			// capture panic, remove game from maps
			if r := recover(); r != nil {
				g.running = false
				log.Printf("Reocvered panic from Game %v listen: %v", id, r)
			}
		}()
		g.listen()
	}()

	s.l.Lock()
	s.cancels[id] = cancel
	s.games[id] = g
	s.l.Unlock()

	return &endless.GameCreated{Code: id}, nil
}

// State ...
func (s *Srv) State(stream endless.Game_StateServer) error {
	var game *Game
	var id string
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Printf("Unable to receive state message during handshake: %v", err)
			return fmt.Errorf("unable to receive message? %v", err)
		}

		r := msg.GetRegister()
		if r == nil {
			out := &endless.Output{
				Data: &endless.Output_Msg{
					Msg: &endless.EventMessage{
						Msg:     "You must register your client first",
						IsError: true,
					},
				},
			}
			stream.Send(out)
			continue
		}

		c := r.GetCode()
		g, ok := s.games[c]
		if !ok {
			out := &endless.Output{
				Data: &endless.Output_Msg{
					Msg: &endless.EventMessage{
						Msg:     fmt.Sprintf("No game with code '%v'", c),
						IsError: true,
					},
				},
			}
			stream.Send(out)
			continue
		}
		game = g
		id = strings.TrimSpace(r.GetId())
		break
	}

	return game.registerClient(id, stream)
}
