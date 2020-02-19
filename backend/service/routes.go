package service

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/seanhagen/endless_stream/backend/endless"
	"github.com/seanhagen/endless_stream/backend/game"
)

// Create ...
func (s *Srv) Create(origCtx context.Context, in *endless.CreateGame) (*endless.GameCreated, error) {
	var id string

	for {
		id = game.GetGameId()
		if _, ok := s.games[id]; ok {
			continue
		}
		break
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	g, err := game.Create(ctx, id, s.entityCol)
	if err != nil {
		cancel()
		return nil, err
	}
	go func() {
		defer func() {
			// capture panic, remove game from maps
			if r := recover(); r != nil {
				g.Running = false
				log.Printf("Reocvered panic from Game %v listen: %v", id, r)
			}
		}()
		g.Listen()
	}()

	s.l.Lock()
	s.cancels[id] = cancel
	s.games[id] = g
	s.l.Unlock()

	return &endless.GameCreated{Code: id}, nil
}

// State ...
func (s *Srv) State(stream endless.Game_StateServer) error {
	var g *game.Game
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
		var ok bool
		g, ok = s.games[c]
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

		id = strings.TrimSpace(r.GetId())
		name = r.GetName()
		break
	}

	if id == "" {
		x, _ := uuid.NewV4()
		id = x.String()
	}

	if name == "" {
		name = fmt.Sprintf("Anon-%v", strings.ToLower(game.GetGameId()))
	}

	log.Printf("registering client %v", id)
	err := g.RegisterClient(id, name, stream)
	log.Printf("client %v done", id)
	return err
}
