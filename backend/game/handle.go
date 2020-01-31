package game

import (
	"context"

	"github.com/seanhagen/endless_stream/backend/endless"
)

// handleInput ...
func (g *Game) handleInput(ctx context.Context, in input) error {
	if in.isPlayer {
		return g.handlePlayerInput(ctx, in.in)
	}

	return g.handleAudienceInput(ctx, in.in)
}

// handleAudienceInput ...
func (g *Game) handleAudienceInput(ctx context.Context, in *endless.Input) error {
	return nil
}

// handlePlayerInput ...
func (g *Game) handlePlayerInput(ctx context.Context, in *endless.Input) error {
	return nil
}
