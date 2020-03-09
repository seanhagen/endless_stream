package game

import (
	"github.com/seanhagen/endless_stream/backend/endless"
	luar "layeh.com/gopher-luar"
)

// how this all works:
//
// example: player using basic attack on an enemy
//   - player sends 'UseSkill' input
//   - player object takes input, creates actionMessage and stores it within itself
//   - on tick, game calls `act()` on the player struct, which returns the actionMessage
//   - game validates that the actionMessage is valid ( targets exist and can be targetted,
//      player has enough focus to use skill, a valid actionType is retured, etc )
//   - if skill is not valid, outputs an error message
//   - if the skill is valid:
//     - pass the actionMessage to each target by calling `apply()` on each actor targetted
//     - the actor calls the `apply()` method of the message by passing in it's own creature

type actionType int32

const (
	action_free  actionType = 0
	action_basic            = 1
)

type actionResult struct {
	harm     int32
	heal     int32
	accuracy int32
}

// actor is something that can act or be acted upon
type actor interface {
	ID(l *luar.LState) int
	// health returns the current and max health of the entity
	Health() (int32, int32)
	// type returns the type of the actor
	Type() endless.Type

	// apply takes the following:
	//   - the creature causing the action
	//   - the action itself
	//   - the current game state
	apply(*creature, actionMessage, *Game) error
	act(*waveState) actionMessage

	// tick is called on every tick
	tick() (*endless.EventMessage, error)
	// round is called at the start of every round, and should do things like:
	//  - run any status scripts ( apply poison effect, etc )
	round() (*endless.EventMessage, error)

	// initiative is called at the start of a wave to determine initiative order.
	// lower is better ( initiative counts 0->20 )
	initiative() int

	// tells the entity it has to take damage
	takeDamage(amount, accuracy int32) *endless.EventMessage
}
