package game

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/seanhagen/endless_stream/backend/endless"
)

// tick ...
func (g *Game) tick(ctx context.Context, t time.Time) error {
	g.Lock()
	defer g.Unlock()
	defer g.tickOut(t)

	// run any AI scripts

	// get player inputs
	playerInputs := g.getPlayerInput()

	// get audience inputs
	// audienceInputs := g.getAudienceInputs()
	// log.Printf("got %v audience inputs", len(audienceInputs))

	g.handleTickCountdowns(ctx)

	// log.Printf("processing game state")
	switch g.screenState.MustState().(GameState) {
	case StateCharSelect:
		g.stateCharSelect(playerInputs)

	case StateNewWave:
		err := g.setupNewWave(ctx)
		if err != nil {
			log.Printf("unable to setup next wave: %v", err)
			return err
		}
	case StateWave:
	// check
	//  if players are all dead, go to 'Defeat' state
	//  if monsters are all dead, go to 'Victory' state
	//  otherwise
	//
	case StateWaveInput:
		if err := g.getCurrentPlayerInput(); err != nil {
			log.Printf("error getting input of current player: %v", err)
			return err
		}
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
	// log.Printf("tick over")

	return nil
}

// tickOut ...
func (g *Game) tickOut(t time.Time) {
	ts, _ := ptypes.TimestampProto(t)
	g.output <- &endless.Output{
		Data: &endless.Output_Tick{
			Tick: &endless.Tick{
				Time:     ts,
				Progress: g.outputCountdowns,
			},
		},
	}
}

// handleTickCountdowns ...
func (g *Game) handleTickCountdowns(ctx context.Context) {
	idx := 0
	for _, tk := range g.tickCountdowns {
		tk.counter--
		if tk.counter == 0 {
			tk.fn(ctx)
			tk = nil
			idx++
		}
	}
	g.tickCountdowns = g.tickCountdowns[:idx]
}
