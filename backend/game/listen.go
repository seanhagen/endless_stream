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

		select {
		case newClient := <-g.newClients:
			g.lock.Lock()
			log.Printf("client connected: %v", newClient.id)
			if newClient.isPlayer {
				g.players[newClient] = true
			} else {
				g.audience[newClient] = true
			}
			g.lock.Unlock()
			// output state attempts to gain lock
			// log.Printf("generating output to send")
			// newClient.out <- g.outputState()
			log.Printf("got new client")

		case clientLeft := <-g.closingClients:
			log.Printf("client disconnected: %v", clientLeft.id)
			g.unregisterHuman(clientLeft)
		default:
		}

		select {
		case update := <-g.output:
			if len(g.players) > 0 {
				// log.Printf("sending output to players")
				for c := range g.players {
					// log.Printf("player '%v' (isPlayer: %v)", c.id, c.isPlayer)
					c.out <- update
				}
			}

		case input := <-g.input:
			log.Printf("got player/audience input: %v", input)
			g.handleInput(input)
		default:
		}

		select {
		case <-stateTick.C:
			g.output <- g.outputState()

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
