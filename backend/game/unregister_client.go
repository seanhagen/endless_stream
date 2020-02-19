package game

import "log"

// unregisterHuman is called in listen.go when a clien disconnects
func (g *Game) unregisterHuman(o output) error {
	g.lock.Lock()
	defer g.lock.Unlock()
	log.Printf("unregister human")

	if o.isPlayer {
		delete(g.players, o)
		g.playerIds[o.id]--
	} else {
		delete(g.audience, o)
	}

	return nil
}
