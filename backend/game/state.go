package game

import (
	"github.com/seanhagen/endless_stream/backend/endless"
)

type skill interface{}

type item interface{}

type status interface {
	effect(*creature)
	tick()
	shouldRemove() bool
}

type creature struct {
	id   string
	name string

	position *endless.Position
	statuses []status

	strength         int32
	maxVitality      int32 // aka max hp
	currentVitality  int32 // aka current hp
	combatDamageBase int32
	vitalityRegen    int32

	intelligence int32
	currentFocus int32
	maxFocus     int32
	willpower    int32
	focusRegen   int32

	agility    int32
	evasion    int32
	accuracy   int32
	initiative int32

	// a list of current modifiers that affect the various stats
	modifiers map[string]int32

	gold int32
	xp   int32
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

type actor interface {
	isPlayer() bool
	id() string
	setNextAction()
}

// outputState ...
func (g *Game) outputState() *endless.CurrentState {

	return &endless.CurrentState{
		Display: g.display,
		// Wave: g.waveNumber,
	}
}
