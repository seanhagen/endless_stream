package game

import (
	"github.com/seanhagen/endless_stream/backend/endless"
)

// output ...
func (g *Game) sendOutput(o *endless.Output) {
	// g.Lock()
	// defer g.Unlock()

	if len(g.players) > 0 {
		for c, b := range g.players {
			if b {
				c.out <- o
			}
		}
	}

	if len(g.audience) > 0 {
		for a, b := range g.audience {
			if b {
				a.out <- o
			}
		}
	}

	if len(g.displayClients) > 0 {
		for d, b := range g.displayClients {
			if b {
				d.out <- o
			}
		}
	}
}
