package game

import (
	"context"
	"log"
	"time"
)

// maxIdle is how long a game can have no connected clients before closing
const maxIdle = 300

// IsRunning ...
func (g *Game) IsRunning(t time.Time) bool {
	if !g.running {
		return false
	}

	if len(g.players) > 0 {
		g.idleTime = 0
		g.idleSince = time.Now()
		return true
	}

	if g.idleTime == 0 {
		g.idleSince = t
		g.idleTime++
		return true
	}

	g.idleTime++
	d := time.Since(g.idleSince)

	if g.idleSince.Before(t) && d.Truncate(time.Second).Seconds() >= maxIdle {
		log.Printf("game has been idle too long ( %v seconds )", d.Seconds())
		return false
	}
	log.Printf("game has been idle since %v (%v seconds)", g.idleSince, d.Seconds())

	// if g.idleTime > maxIdle {
	// 	return false
	// }

	return true
}

// Shutdown ...
func (g *Game) Shutdown(ctx context.Context) error {
	g.cancelFn()
	g.running = false
	return nil
}
