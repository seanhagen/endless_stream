package game

import (
	"log"

	"github.com/seanhagen/endless_stream/backend/endless"
)

type player struct {
	creature

	class endless.Class

	isAI  bool
	level int32

	skills    map[string]skill
	inventory map[string]item
}

type action interface {
	isItem() bool
	isSkill() bool
	isSkip() bool
	do(*Game)
}

type actor interface {
	// tick is called every tick
	tick() error
	// round is called at the start of every round
	round() error
	// getAction is called when it's this actors turn in iniative order
	getAction() action
	// iniative determines iniative order, lower goes earlier in a round
	initiative() int

	OnHit()
	OnMove()
	OnDeath()
}

// outputState ...
func (g *Game) outputState() *endless.Output {
	g.Lock()
	defer g.Unlock()
	log.Printf("outputting state")

	var d endless.Display
	switch g.screenState.MustState().(GameState) {
	case StateCharSelect:
		d = endless.Display_ScreenCharSelect
	case StateNewWave:
		d = endless.Display_ScreenNewWave
	case StateWaveAnimWait:
		fallthrough
	case StateWaveInput:
		fallthrough
	case StateWaveProcess:
		fallthrough
	case StateWave:
		d = endless.Display_ScreenWave
	case StateDefeat:
		d = endless.Display_ScreenDead
	case StateGameOver:
		d = endless.Display_ScreenGameOver
	case StateVictory:
		d = endless.Display_ScreenVictory
	case StateStore:
		d = endless.Display_ScreenStore
	default:
		d = endless.Display_ScreenLoading
	}

	sc := map[string]endless.ClassType{}
	if len(g.selectedCharacters) > 0 {
		for k, v := range g.selectedCharacters {
			sc[v] = k
		}
	}

	return &endless.Output{
		Data: &endless.Output_State{
			State: &endless.CurrentState{
				Display:  d,
				Selected: &endless.CharacterSelected{Selected: sc},
				// Wave: g.waveNumber,
			},
		},
	}
}
