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

type actor interface {
	hasScript() bool
	runScript(*gameState) error
	setNextAction(action)
	action(*gameState) action
}

type action interface {
	perform()
}

type gameState struct {
	code string

	current_initiative_step int
	current_initiative      int
	max_initiative          int

	players  map[string]player
	audience map[string]audience
	monsters map[string]monster

	initiative map[int][]actor
}

func newState(ctx context.Context, id string) *gameState {
	return &gameState{
		code:                    id,
		current_initiative_step: 0,
		current_initiative:      0,
		max_initiative:          0,
		players:                 map[string]player{},
		audience:                map[string]audience{},
		monsters:                map[string]monster{},
	}
}

// tick ...
func (gs *gameState) tick(t time.Time) error {
	// get current actor

	//

	// actor has action to perform?

	// if no, return

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
