package game

import "log"

// unregisterHuman is called in listen.go when a clien disconnects
func (g *Game) unregisterHuman(o output) error {
	log.Printf("client disconnected: %v, unregistering human", o.id)
	g.Lock()
	defer g.Unlock()

	g.connected--
	if o.isPlayer {
		close(o.out)
		delete(g.players, o)
		g.playerIds[o.id]--
	} else {
		delete(g.audience, o)
	}

	log.Printf("human unregistered")
	return nil
}
