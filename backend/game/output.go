package game

import (
	"log"

	"github.com/seanhagen/endless_stream/backend/endless"
)

// output ...
func (g *Game) sendOutput(o *endless.Output) {
	g.Lock()
	defer g.Unlock()

	log.Printf("####### sending output to %v players", len(g.players))
	if len(g.players) > 0 {
		for c := range g.players {
			log.Printf("player '%v' (isPlayer: %v)", c.id, c.isPlayer)
			c.out <- o
			log.Printf("sent update to %v", c.id)
		}
	}
}
