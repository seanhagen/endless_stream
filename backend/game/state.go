package game

import (
	"github.com/seanhagen/endless_stream/backend/endless"
)

// outputState ...
func (g *Game) outputState() *endless.Output {
	g.Lock()
	defer g.Unlock()
	// log.Printf("outputting state")

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

	// log.Printf("selected: %v", spew.Sdump(sc, g.selectedCharacters))

	return &endless.Output{
		Data: &endless.Output_State{
			State: &endless.CurrentState{
				Display:       d,
				Selected:      &endless.CharacterSelected{Selected: g.selectedCharacters},
				AudienceCount: int64(len(g.audienceIds)),
				//Players: map[string]*endless.Player,
				//Monsters: []*endless.Monster,
				//CurrentPlayer: *wrappers.StringValue, // the uuid of the current player
				//CurrentWave: *endless.Wave,
				//UpcomingWaves: map[int32]*endless.Wave,
			},
		},
	}
}
