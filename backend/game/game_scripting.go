package game

import (
	"log"

	lua "github.com/yuin/gopher-lua"
)

// setupFunctions attaches all the game state functions to a script.
//
// The idea is that before a skill script or a AI script is run, we want to
// register all of the game functions so that the scripts can do what they
// need to.
//
// Functions to register:
//  - spawnMonster -- given a name ( or id? ), will add the given monster to the current wave
//  - haveMemory -- given a key, will check if that is registered in the game's "memory"
//  - outputMsg -- allow a script to output a message
func (g *Game) setupFunctions(l *lua.LState) error {
	l.SetGlobal("spawnMonster", l.NewFunction(g.spawnMonster))
	l.SetGlobal("haveMemory", l.NewFunction(g.haveMemory))
	l.SetGlobal("outputMsg", l.NewFunction(g.outputMsg))
	return nil
}

// spawnMonster ...
func (g *Game) spawnMonster(l *lua.LState) int {
	id := l.ToString(1)
	log.Printf("script wants to spawn monster: %v", id)
	return 0
}

// haveMemory ...
func (g *Game) haveMemory(l *lua.LState) int {
	key := l.ToString(1)
	_, ok := g.Memory[key]
	l.Push(lua.LBool(ok))
	return 1
}

// outputMsg is like a `fmt.Sprintf`, with some custom template variables
func (g *Game) outputMsg(l *lua.LState) int {
	msg := l.ToString(1)
	log.Printf("script wants to output a message: %v", msg)
	return 0
}
