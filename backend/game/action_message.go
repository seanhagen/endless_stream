package game

import (
	"github.com/seanhagen/endless_stream/backend/endless"
)

type actionMessage interface {
	// targets returns a slice containing the ids of what this action
	targets() []string
	// cost is the focus cost of the action, used by the game to determine if the
	// action is valid.
	//
	// the second return value is the action type, which indicates if the action is 'free'
	// or 'basic' -- free actions don't use up the actors turn ( ie, using a health potion
	// is a free action, and so would allow the actor to do something else their turn )
	cost() (int32, actionType)

	// apply takes the creature that created the action, and the creature it's being applied to,
	// as well as the current game state and does the thing
	apply(from, to *creature, g *Game) error

	// if the action requires the game to output a message
	output() *endless.EventMessage
}

var _ actionMessage = &skipMsg{}
var _ actionMessage = &move{}

type move struct {
	cid string
	dir string
}

// targets ...
func (m move) targets() []string {
	return []string{m.cid}
}

// cost ...
func (m move) cost() (int32, actionType) {
	return 0, action_basic
}

// apply ...
func (m move) apply(_, to *creature, _ *Game) error {
	if m.dir == "left" {
		m.moveLeft(to)
	} else {
		m.moveRight(to)
	}
	return nil
}

// output ...
func (m move) output() *endless.EventMessage {
	// TODO: output a message when a creature moves
	return nil
}

// moveRight ...
func (m move) moveRight(cr *creature) {
	switch cr.Position {
	case endless.Position_Left:
		cr.Position = endless.Position_Middle
	case endless.Position_Middle:
		cr.Position = endless.Position_Right
	}
}

// moveLeft ...
func (m move) moveLeft(cr *creature) {
	switch cr.Position {
	case endless.Position_Middle:
		cr.Position = endless.Position_Left
	case endless.Position_Right:
		cr.Position = endless.Position_Middle
	}
}

// skipMsg is a message returned when
type skipMsg struct {
	reason string
}

func (s skipMsg) targets() []string {
	return []string{}
}

// cost for a skip action is 0, but the action is basic -- meaning it won't use focus
// but will cause the actor to lose their turn
func (s skipMsg) cost() (int32, actionType) {
	return 0, action_basic
}

// apply ...
func (s skipMsg) apply(_, _ *creature, _ *Game) error {
	return nil
}

// output ...
func (s skipMsg) output() *endless.EventMessage {
	if s.reason == "" {
		return nil
	}
	return &endless.EventMessage{
		Msg: s.reason,
	}
}

var _ actionMessage = &basicAttack{}

type basicAttack struct {
	target   string
	damage   int32
	accuracy int32
	cst      int32
}

// targets ...
func (ba basicAttack) targets() []string {
	return []string{ba.target}
}

// cost ...
func (ba basicAttack) cost() (int32, actionType) {
	return ba.cst, action_basic
}

// apply ...
func (ba basicAttack) apply(from, to *creature, g *Game) error {
	from.CurrentFocus -= ba.cst
	to.takeDamage(ba.damage, ba.accuracy)
	return nil
}

// output ...
func (ba basicAttack) output() *endless.EventMessage {
	return nil
}
