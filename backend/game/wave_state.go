package game

import "github.com/seanhagen/endless_stream/backend/endless"

type waveState struct {
	current_initiative_step int
	current_initiative      int
	max_initiative          int

	initiative map[int][]actor

	info endless.Wave

	// monsterData is a map of string -> data that is passed in when a creature
	// performs various actions
	//
	// For example, when a cultist dies it increments a counter -- when the counter hits 7, a shoggoth is summoned.
	//
	// That information is stored here.
	monsterData map[string]interface{}
}

// current ...
func (ws waveState) current() actor {
	cs := ws.current_initiative_step
	return ws.initiative[ws.current_initiative][cs]
}
