package game

import (
	"context"
	"time"

	"github.com/seanhagen/endless_stream/backend/endless"
)

type human interface {
}

type player struct{}

type audience struct{}

type monster interface {
	runScript(*gameState) error
}

type creature interface{}

type gameState struct {
	code string

	current_initiative int

	players  map[string]player
	audience map[string]audience
	monsters map[string]monster
}

func newState(ctx context.Context, id string) *gameState {
	return &gameState{
		code:               id,
		current_initiative: 0,
		players:            map[string]player{},
		audience:           map[string]audience{},
		monsters:           map[string]monster{},
	}
}

// tick ...
func (gs *gameState) tick(t time.Time) error {
	// log.Printf("game state %v ticking onwards", gs.code)
	return nil
}

// handleInput ...
func (gs *gameState) handleInput(ctx context.Context, in input) error {
	if in.isPlayer {
		return gs.handlePlayerInput(ctx, in.in)
	}

	return gs.handleAudienceInput(ctx, in.in)
}

// handleAudienceInput ...
func (gs *gameState) handleAudienceInput(ctx context.Context, in *endless.Input) error {
	return nil
}

// handlePlayerInput ...
func (gs *gameState) handlePlayerInput(ctx context.Context, in *endless.Input) error {
	return nil
}
