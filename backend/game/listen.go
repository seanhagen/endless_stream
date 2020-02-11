package game

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/seanhagen/endless_stream/backend/endless"
)

// listen ...
func (g *Game) listen() {
	ticker := time.NewTicker(tickLen)
	stateTick := time.NewTicker(stateLen)

	for {
		var err error
		select {
		case newClient := <-g.newClients:
			log.Printf("client connected: %v", newClient.id)
			g.lock.Lock()
			g.players[newClient] = true
			g.lock.Unlock()

		case clientLeft := <-g.closingClients:
			log.Printf("client disconnected: %v", clientLeft.id)
			g.unregisterHuman(clientLeft)

		case update := <-g.output:
			if len(g.players) > 0 {
				// log.Printf("sending output to players")
				for c := range g.players {
					// log.Printf("player '%v' (isPlayer: %v)", c.id, c.isPlayer)
					c.out <- update
				}
			}

		case input := <-g.input:
			log.Printf("got player input: %v", input)
			ctx, cancel := context.WithTimeout(g.ctx, time.Second)
			err = g.handleInput(ctx, input)
			cancel()

		case t := <-ticker.C:
			// log.Printf("game tick")
			err := g.tick(t)
			if err != nil {
				log.Printf("unable to to tick: %v", err)
			}

			ts, _ := ptypes.TimestampProto(t)
			g.output <- &endless.Output{
				Data: &endless.Output_Tick{
					Tick: &endless.Tick{Time: ts},
				},
			}

		case <-stateTick.C:
			// log.Printf("sending state")
			g.output <- &endless.Output{
				Data: &endless.Output_State{
					State: &endless.CurrentState{},
				},
			}

		case <-g.ctx.Done():
			log.Printf("game context signaled done!")
			goto finished
		}

		if err != nil {
			log.Printf("Error occured: %v", err)
		}
	}
finished:
	log.Printf("game done")
}
