package game

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/seanhagen/endless_stream/backend/endless"
)

// handleInput ...
func (g *Game) handleInput(ctx context.Context, in input) error {
	err := g.state.handleInput(ctx, in)

	if err != nil {
		g.lock.Lock()
		defer g.lock.Unlock()
		mid := g.msgId
		g.msgId++

		g.output <- &endless.Output{
			Data: &endless.Output_Msg{
				Msg: &endless.EventMessage{
					MsgId: mid,
					// TODO: error -> json string? need error messages that make sense to the players?
					Msg:      fmt.Sprintf("Unable to handle input: %v", err),
					IsError:  true,
					PlayerId: &wrappers.StringValue{Value: in.in.GetPlayerId()},
				},
			},
		}
	}
	return err
}
