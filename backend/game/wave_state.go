package game

import (
	"fmt"

	"github.com/seanhagen/endless_stream/backend/endless"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
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

// addActor ...
func (ws waveState) addActor(a actor) error {
	id := a.ID()
	ws.Entities[id] = a
	return nil
}

// EntityByKey ...
func (ws waveState) EntityByKey(l *lua.LState) int {
	id := l.ToString(1)
	if e, ok := ws.Entities[id]; ok {
		o := luar.New(l, e)
		l.Push(o)
		return 1
	}

	return 0
}

// EntityKeys ...
func (ws waveState) EntityKeys(l *lua.LState) int {
	out := l.NewTable()
	for k, _ := range ws.Entities {
		out.Append(lua.LString(k))
	}

	l.Push(out)
	return 1
}

// getPlayers ...
func (ws waveState) getPlayers(l *lua.LState) int {

	return 0
}

// getMonsters ...
func (ws waveState) getMonsters(l *lua.LState) int {

	return 0
}

// register ...
func (ws waveState) register(cr *creature) {
	ek := cr.ls.NewFunction(ws.EntityKeys)
	ebk := cr.ls.NewFunction(ws.EntityByKey)
	gp := cr.ls.NewFunction(ws.getPlayers)
	gm := cr.ls.NewFunction(ws.getMonsters)

	cr.ls.SetGlobal("entityKeys", ek)
	cr.ls.SetGlobal("entityByKey", ebk)
	cr.ls.SetGlobal("getPlayers", gp)
	cr.ls.SetGlobal("getMonsters", gm)
	cr.ls.SetGlobal("waveState", luar.New(cr.ls, ws))
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

	if i, ok := ws.initiative[ws.current_initiative]; ok {
		return i[cs]
	}

	return nil //ws.initiative[ws.current_initiative][cs]
}

// tick ...
func (ws *waveState) tick() error {
	// get current actor
	actr := ws.current()
	if actr == nil {
		return fmt.Errorf("no current actor")
	}

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

// proceed ...
func (ws waveState) proceed() bool {
	return ws.currentAction != nil
}

// process ...
func (ws waveState) process() error {
	return fmt.Errorf("not yet")
}

// roundOver ...
func (ws waveState) waveComplete() bool {
	return true
}

// waveFailed ...
func (ws waveState) waveFailed() bool {
	return true
}
