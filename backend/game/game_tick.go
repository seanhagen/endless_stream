package game

import (
	"context"
	"log"
	"time"
)

// tick ...
func (g *Game) tick(ctx context.Context, t time.Time) error {
	g.Lock()
	defer g.Unlock()
	log.Printf("game tick happens")

	if len(g.players) == 0 {
		g.idleTime++
	}

	// run all the ai scripts

	// get player inputs
	playerInputs := map[string][]input{}
	l := len(g.playerInput)
	log.Printf("player inputs to be processed: %v", l)
	if l > 0 {
		for i := 0; i < l; i++ {
			pi := <-g.playerInput
			id := pi.in.GetPlayerId()
			ins, ok := playerInputs[id]
			if !ok {
				ins = []input{}
			}
			ins = append(ins, pi)
			playerInputs[id] = ins
		}
	}

	// get audience inputs
	audienceInputs := []input{}
	l = len(g.audienceInput)
	log.Printf("audience inputs: %v", l)
	for i := 0; i < l; i++ {
		ai := <-g.audienceInput
		if y := ai.in.GetAudience(); y != nil {
			audienceInputs = append(audienceInputs, ai)
		}
	}

	// check tickCountdown timers
	cds := g.tickCountdowns[g.tickCounterIdx]
	log.Printf("countdowns to process: %v", len(cds))
	for _, c := range cds {
		c(ctx)
	}
	// remove those counters
	g.tickCountdowns[g.tickCounterIdx] = []countdownFunc{}
	// decrement counter
	g.tickCounterIdx--
	if g.tickCounterIdx < tickCounterMin { // loop back around if we've hit min ( ie, 0 )
		g.tickCounterIdx = tickCounterMax
	}

	log.Printf("processing game state")
	switch g.screenState.MustState().(GameState) {
	case StateCharSelect:
		g.stateCharSelect(playerInputs)

	case StateNewWave:
		// if player hasn't selected a class, assign random from classes left over
		// if there are fewer than 4 players, create AI players to fill the slots
		// set up next wave
		// run all round counters
	case StateWave:
	// check
	//  if players are all dead, go to 'Defeat' state
	//  if monsters are all dead, go to 'Victory' state
	//  otherwise
	//
	case StateWaveInput:
		// get current actor
		//   get input
		//   if valid, store and continue
	case StateWaveProcess:
		// process current actor action
		//   if action is move or skill, advance iniative
		//   otherwise apply item affect and continue to next state
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
	log.Printf("tick over")
	return nil
}
