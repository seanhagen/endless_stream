package game

import (
	"github.com/seanhagen/endless_stream/backend/endless"
)

var _ actor = &player{}

type player struct {
	*creature

	class     endless.ClassType
	inventory inventory

	nextAction actionMessage

	is_ai bool
}

// Entity interface methods
// round, initiative, & health are covered by creature
// health is covered by creature struct
// tick ...
func (p *player) tick() (*endless.EventMessage, error) {
	return nil, nil
}

// setAction ...
func (p *player) setAction(inp *endless.Input) {
	var next actionMessage
	switch v := inp.GetInput().(type) {
	case *endless.Input_Skill:
		next = p.setActionSkill(v)

	case *endless.Input_Item:
		next = p.setActionItem(v)

	case *endless.Input_Move:
		next = p.setActionMove(v)
	}
	p.nextAction = next
}

// Actor interface methods
// apply ...
func (p *player) apply(from *creature, am actionMessage, g *Game) error {
	return am.apply(from, p.creature, g)
}

// act ...
func (p *player) act(ws *waveState) actionMessage {
	if p.is_ai {
		return p.creature.act(ws)
	}
	return p.nextAction
}
