package game

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/seanhagen/endless_stream/backend/endless"
)

// Create ...
func (s *Srv) Create(origCtx context.Context, in *endless.CreateGame) (*endless.GameCreated, error) {
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
	var id, name string
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
		name = r.GetName()
		break
	}

	if id == "" {
		x, _ := uuid.NewV4()
		id = x.String()
	}

	if name == "" {
		name = fmt.Sprintf("Anon-%v", strings.ToLower(getGameId()))
	}

	return game.registerClient(id, name, stream)
}
