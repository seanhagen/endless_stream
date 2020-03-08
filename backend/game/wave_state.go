package game

import (
	"github.com/seanhagen/endless_stream/backend/endless"
	lua "github.com/yuin/gopher-lua"
)

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
	MonsterData map[string]interface{}

	Entities map[string]actor

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
		MonsterData:             map[string]interface{}{},
		Entities:                map[string]actor{},
	}
}

// EntityKeys ...
func (ws waveState) EntityKeys(l *lua.LState) int {
	//out := []string{}

	out := l.NewTable()
	for k, _ := range ws.Entities {
		//out = append(out, k)
		out.Append(lua.LString(k))
	}

	// t := l.ToTable(out)
	// l.Push(t)

	//lua.LTable

	l.Push(out)

	//return out

	return 1
}

// waveStart ...
func (ws *waveState) waveStart() error {
	for _, actr := range ws.Entities {
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
func (ws *waveState) tick() error {
	// get current actor
	actr := ws.current()

	// get input
	if act := actr.act(ws); act != nil {
		// check that all targets are entities that are currently in the wave
		valid := true
		for _, t := range act.targets() {
			if _, ok := ws.Entities[t]; !ok {
				valid = false
			}
		}

		// if valid, store and continue
		if valid {
			ws.currentAction = act
		}
	}

	return nil
}

// getFns ...
func (ws waveState) getFns() {

}

// proceed ...
func (ws waveState) proceed() bool {
	return ws.currentAction != nil
}
