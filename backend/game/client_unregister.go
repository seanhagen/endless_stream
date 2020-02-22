package game

import "log"

// unregisterHuman is called in listen.go when a clien disconnects
func (g *Game) unregisterHuman(o output) error {
	log.Printf("client disconnected: %v, unregistering human", o.id)
	g.Lock()
	defer g.Unlock()

	g.connected--
	close(o.out)
	if o.isPlayer {
		delete(g.players, o)
		g.playerIds[o.id]--
	}

	if o.isAudience {
		delete(g.audience, o)
		delete(g.audienceIds, o.id)
	}

	if o.isDisplay {
		delete(g.displayClients, o)
	}

	log.Printf("human unregistered")
	return nil
}
