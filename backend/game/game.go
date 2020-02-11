package game

import (
	"log"
	"time"

	"github.com/gofrs/uuid"
	"github.com/seanhagen/endless_stream/backend/endless"
)

// tick ...
func (g *Game) tick(t time.Time) error {
	g.lock.Lock()
	defer g.lock.Unlock()

	if len(g.players) == 0 {
		g.idleTime++
	}

	err := g.state.tick(t)
	if err != nil {
		return err
	}

	return nil
}

// unregisterHuman ...
func (g *Game) unregisterHuman(o output) error {
	g.lock.Lock()
	defer g.lock.Unlock()

	if o.isPlayer {
		delete(g.players, o)
		g.playerIds[o.id]--
	} else {
		delete(g.audience, o)
	}

	return nil
}

// registerHuman ...
func (g *Game) registerHuman(id, name string) (*endless.Output, output, error) {
	if id == "" {
		x, err := uuid.NewV4()
		if err != nil {
			return nil, output{}, err
		}
		id = x.String()
	}

	// accessing some maps, gotta lock
	g.lock.Lock()
	defer g.lock.Unlock()

	out := output{
		id:  id,
		out: make(chan *endless.Output),
	}

	v, ok := g.playerIds[id]
	if ok && v < 1 {
		// player is rejoining
		log.Printf("player is reconnecting")
	}

	if len(g.players) <= 4 {
		msg, err := g.registerPlayer(id, name)

		out.isPlayer = true
		return msg, out, err
	}

	// g.audienceIds[id] = 1
	msg, err := g.registerAudience(id)
	return msg, out, err
}

// registerPlayer ...
func (g *Game) registerPlayer(id, name string) (*endless.Output, error) {
	g.playerIds[id] = 1
	g.playerCharacters[id] = nil
	g.playerNames[id] = name

	isVip := false
	if len(g.players) == 0 {
		isVip = true
		g.vipPlayer = id
	}

	out := &endless.Output{
		Data: &endless.Output_Joined{
			Joined: &endless.JoinedGame{
				Id:         id,
				AsAudience: false,
				IsVip:      isVip,
				Name:       name,
			},
		},
	}

	return out, nil
}

// registerAudience ...
func (g *Game) registerAudience(id string) (*endless.Output, error) {
	out := &endless.Output{
		Data: &endless.Output_Joined{
			Joined: &endless.JoinedGame{
				Id:         id,
				AsAudience: true,
				Name:       "Audience Member",
			},
		},
	}

	return out, nil
}
