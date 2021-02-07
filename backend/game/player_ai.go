package game

import "github.com/seanhagen/endless_stream/backend/endless"

type player_ai struct {
	*creature

	class      endless.ClassType
	skills     charSkillMap
	inventory  inventory
	nextAction actionMessage
}

// setup ...
func (ai *player_ai) setup(g *Game) error {
	cr, err := ai.creature.spawn(g)
	if err != nil {
		return err
	}
	ai.creature = cr
	return nil
}

// // act ...
// func (ai *player_ai) act(ws *waveState) actionMessage {
// 	// this should really run some ai script to determine what to do
// 	return ai.nextAction
// }

// // tick ...
// func (ai *player_ai) tick() (*endless.EventMessage, error) {
// 	return nil, nil
// }

// // round ...
// func (ai *player_ai) round() (*endless.EventMessage, error) {
// 	return nil, nil
// }

// id(), initiative(), health(), takeDamage() covered by creature
