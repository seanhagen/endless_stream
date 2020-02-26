package game

import "github.com/seanhagen/endless_stream/backend/endless"

type player struct {
	creature

	class endless.Class

	isAI  bool
	level int32

	skills    map[string]skill
	inventory map[string]item
}

/*
type actor interface {
  // tick is called every tick
  tick() error
  // round is called at the start of every round
  round() error
  // getAction is called when it's this actors turn in iniative order
  getAction() action
  // iniative determines iniative order, lower goes earlier in a round
  initiative() int
}

*/

// tick ...
func (p *player) tick() error {
	if p.isAI {
		return p.creature.tick()
	}
	return nil
}

// getAction ...
func (p *player) getAction(inp *endless.Input) action {
	return nil
}

// round ...
func (p *player) round() error {
	return nil
}

// initiative ...
func (p *player) initiative() int {
	return p.creature.iniative()
}
