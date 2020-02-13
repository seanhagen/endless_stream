package game

import (
	"github.com/seanhagen/endless_stream/backend/endless"
)

type skill interface{}

type item interface{}

type creature struct {
	Id   string
	Name string

	Position *endless.Position
	Statuses []status

	Strength         int32
	MaxVitality      int32 // aka max hp
	CurrentVitality  int32 // aka current hp
	CombatDamageBase int32
	VitalityRegen    int32

	Intelligence int32
	CurrentFocus int32
	MaxFocus     int32
	Willpower    int32
	FocusRegen   int32

	Agility    int32
	Evasion    int32
	Accuracy   int32
	Initiative int32

	// a list of current modifiers that affect the various stats
	Modifiers map[string]int32

	Gold int32
	XP   int32
}

type player struct {
	creature

	class endless.Class

	isAI  bool
	level int32

	skills    map[string]skill
	inventory map[string]item
}

type monster struct {
	creature
	mType    endless.Type
	isFlying bool
	isBoss   bool
	cost     int32
}

type action interface{}

type actor interface{}

// outputState ...
func (g *Game) outputState() *endless.Output {
	g.lock.Lock()
	defer g.lock.Unlock()

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

	return &endless.Output{
		Data: &endless.Output_State{
			State: &endless.CurrentState{
				Display: d,
				// Wave: g.waveNumber,
			},
		},
	}
}
