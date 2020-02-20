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
				g.Shutdown(context.Background())
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

	msg, err := stream.Recv()
	if err != nil {
		return fmt.Errorf("unable to receive message? %v", err)
	}

	r := msg.GetRegister()
	if r == nil {
		return fmt.Errorf("First message must be registration message")
	}

	c := r.GetCode()
	var ok bool
	g, ok = s.games[c]
	if !ok {
		return fmt.Errorf("no game with code '%v'", c)
	}

	id = strings.TrimSpace(r.GetId())
	name = r.GetName()

	if id == "" {
		x, _ := uuid.NewV4()
		id = x.String()
	}

	if name == "" {
		name = fmt.Sprintf("Anon-%v", strings.ToLower(game.GetGameId()))
	}

	err = g.RegisterClient(id, name, stream)
	return err
}
