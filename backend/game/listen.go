package game

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/seanhagen/endless_stream/backend/endless"
)

// Listen ...
func (g *Game) Listen() {
	ticker := time.NewTicker(tickLen)
	stateTick := time.NewTicker(stateLen)

	for {
		var err error
		// multiple select statements ( with default cases ) allows go to do more
		// than one thing each iteration, so a tick won't have to wait because
		// there's an update to send out or a player has connected

		// select {
		// // case newClient := <-g.newClients:
		// // 	log.Printf("got new client")
		// // g.Lock()
		// // if newClient.isPlayer {
		// // 	g.players[newClient] = true
		// // } else {
		// // 	g.audience[newClient] = true
		// // }
		// // g.Unlock()
		// // output state attempts to gain lock
		// // log.Printf("generating output to send")
		// // newClient.out <- g.outputState()

		// // case clientLeft := <-g.closingClients:
		// // 	g.unregisterHuman(clientLeft)
		// default:
		// }

		select {
		case update := <-g.output:
			g.sendOutput(update)
		case input := <-g.input:
			g.handleInput(input)
		case <-stateTick.C:
			g.output <- g.outputState()
		default:
		}

		select {
		case t := <-ticker.C:
			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, time.Second*2)
			err := g.tick(ctx, t)
			if err != nil {
				log.Printf("unable to to tick: %v", err)
			}
			cancel()
			ts, _ := ptypes.TimestampProto(t)
			g.output <- &endless.Output{
				Data: &endless.Output_Tick{
					Tick: &endless.Tick{Time: ts},
				},
			}
		default:
		}

		select {
		case <-g.ctx.Done():
			log.Printf("game context signaled done!")
			goto finished
		default:
		}

		if err != nil {
			log.Printf("Error occured: %v", err)
		}
	}
finished:
	log.Printf("game done")
}

// finished ...
func (g *Game) finished() {
	g.cancelFn()
}
