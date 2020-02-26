package game

import (
	"context"
	"log"
	"time"
)

// Listen ...
func (g *Game) Listen() {
	ticker := time.NewTicker(tickLen)
	stateTick := time.NewTicker(stateLen)

	for {
		// multiple select statements ( with default cases ) allows go to do more
		// than one thing each iteration, so a tick won't have to wait because
		// there's an update to send out or a player has connected

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
		default:
		}

		select {
		case <-g.ctx.Done():
			goto finished
		default:
		}
	}
finished:
	log.Printf("game done")
}

// finished ...
func (g *Game) finished() {
	g.cancelFn()
}
