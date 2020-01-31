package game

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/seanhagen/endless_stream/backend/endless"
)

// listen ...
func (g *Game) listen() {
	ticker := time.NewTicker(tickLen)
	stateTick := time.NewTicker(stateLen)

	msgId := 1

	for {
		select {
		case newClient := <-g.newClients:
			log.Printf("client connected: %v", newClient.id)
			g.players[newClient] = true

		case clientLeft := <-g.closingClients:
			log.Printf("client disconnected: %v", clientLeft.id)
			delete(g.players, clientLeft)

		case update := <-g.output:
			if len(g.players) > 0 {
				log.Printf("sending output to players")
				for c := range g.players {
					log.Printf("player '%v' (isPlayer: %v)", c.id, c.isPlayer)
					c.out <- update
				}
			}

		case input := <-g.input:
			log.Printf("got player input: %v", input)
			ctx, cancel := context.WithTimeout(g.ctx, time.Second)
			err := g.handleInput(ctx, input)
			if err != nil {
				g.output <- &endless.Output{
					Data: &endless.Output_Msg{
						Msg: &endless.EventMessage{
							MsgId: int32(msgId),
							// TODO: error -> json string? need error messages that make sense to the players?
							Msg:      fmt.Sprintf("Unable to handle input: %v", err),
							IsError:  true,
							PlayerId: &wrappers.StringValue{Value: input.in.GetPlayerId()},
						},
					},
				}
				msgId++
			}
			cancel()

		case tick := <-ticker.C:
			log.Printf("game tick")
			g.state.tick(tick)
			ts, _ := ptypes.TimestampProto(tick)

			g.output <- &endless.Output{
				Data: &endless.Output_Tick{
					Tick: &endless.Tick{Time: ts},
				},
			}

		case <-stateTick.C:
			log.Printf("sending state")
			g.output <- &endless.Output{
				Data: &endless.Output_State{
					State: &endless.CurrentState{},
				},
			}

		case <-g.ctx.Done():
			log.Printf("game context signaled done!")
			goto finished
		}
	}
finished:
	log.Printf("game done")
}
