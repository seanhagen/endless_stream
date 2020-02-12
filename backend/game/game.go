package game

import (
	"context"
	"log"
	"time"

	"github.com/gofrs/uuid"
	"github.com/seanhagen/endless_stream/backend/endless"
)

// tick ...
func (g *Game) tick(ctx context.Context, t time.Time) error {
	g.lock.Lock()
	defer g.lock.Unlock()

	if len(g.players) == 0 {
		g.idleTime++
	}

	// run all the ai scripts

	// check tickCountdown timers
	cds := g.tickCountdowns[g.tickCounterIdx]
	for _, c := range cds {
		c(ctx)
	}
	// remove those counters
	g.tickCountdowns[g.tickCounterIdx] = []countdownFunc{}
	// decrement counter
	g.tickCounterIdx--
	if g.tickCounterIdx < tickCounterMin { // loop back around
		g.tickCounterIdx = tickCounterMax
	}

	switch g.screenState.MustState().(GameState) {
	case StateCharSelect:
		// handle assigning classes to players
		// if VIP sends 'GameStart', move to new wave state
	case StateNewWave:
		// set up next wave
		// run all round counters
	case StateWave:
	// check
	//  if players are all dead, go to 'Defeat' state
	//  if monsters are all dead, go to 'Victory' state
	//  otherwise go to wave input state
	case StateWaveInput:
		// get current actor
		//   get input
		//   if valid, store and continue
	case StateWaveProcess:
		// process current actor action
		//   if action is move or skill, advance iniative
		//   otherwise apply item affect and continue
		// send tick to all current status effects
		// check all monsters
		//   if dead, run 'onDeath' script then remove
		// check all players
		//   if dead, send status update to that player
		// determine next in iniative order
		// any other state updates
		// start countdown timer for animation wait
	case StateWaveAnimWait:
		// wait for timer to finish
	case StateDefeat:
		// output defeat message
		// if vip sends 'StartOver', go to character select state
		// if vip sends 'IGiveUp', go to game over state
	case StateVictory:
		// output victory message
		// start countdown timer for transition
		// when timer is done or vip sends continue, go to store state
	case StateStore:
		// send store inventory message
		// start countdown timer for transition
		// when timer is done or vip sends continue, go to new wave
	case StateGameOver:
		// start countdown timer
		// when timer over or vip sends quit message, quit game
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
