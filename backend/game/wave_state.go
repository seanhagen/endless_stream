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

	entities map[string]actor

	currentAction actionMessage

	xpGained   int32
	goldGained int32
}

func newWaveState() *waveState {
	return &waveState{
		current_initiative:      30,
		current_initiative_step: 0,
		max_initiative:          30,
		initiative:              map[int][]actor{},
		monsterData:             map[string]interface{}{},
		entities:                map[string]actor{},
	}
}

// waveStart ...
func (ws *waveState) waveStart() error {
	for _, actr := range ws.entities {
		i := actr.initiative()
		ini, ok := ws.initiative[i]
		if !ok {
			ini = []actor{}
		}
		ini = append(ini, actr)
		ws.initiative[i] = ini
	}
	return nil
}

// current ...
func (ws waveState) current() actor {
	cs := ws.current_initiative_step
	return ws.initiative[ws.current_initiative][cs]
}

// tick ...
func (ws waveState) tick() error {
	// get current actor
	actr := ws.current()

	// get input
	if act := actr.act(); act != nil {
		// if valid, store and continue

		// check that all targets are entities that are currently in the wave
		valid := true
		for _, t := range act.targets() {
			if _, ok := ws.entities[t]; !ok {
				valid = false
			}
		}

		if valid {
			ws.currentAction = act
		}
	}

	return nil
}

// proceed ...
func (ws waveState) proceed() bool {
	return ws.currentAction != nil
}
