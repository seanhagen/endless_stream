package game

import (
	"github.com/seanhagen/endless_stream/backend/endless"
)

var _ actor = &player{}

type player struct {
	*creature

	class     endless.Class
	isAI      bool
	skills    charSkillMap
	inventory inventory

	nextAction actionMessage
}

// Entity interface methods
// round, initiative, & health are covered by creature
// health is covered by creature struct
// tick ...
func (p *player) tick() (*endless.EventMessage, error) {
	if p.isAI {
		return p.creature.tick()
	}
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
func (p *player) apply(am actionMessage, g *Game) error {
	return am.apply(p.creature, g)
}

// act ...
func (p *player) act() actionMessage {
	return p.nextAction
}
