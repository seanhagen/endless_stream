package game

import (
	"bytes"
	"log"

	"github.com/seanhagen/endless_stream/backend/endless"
	lua "github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
	luar "layeh.com/gopher-luar"
)

type creature struct {
	Id          string
	Name        string
	Description string

	Position *endless.Position
	Statuses []Status

	Strength         int32
	MaxVitality      int32 // aka max hp
	CurrentVitality  int32 // aka current hp
	CombatDamageBase int32
	VitalityRegen    int32

	Intelligence int32
	CurrentFocus int32
	MaxFocus     int32
	Willpower    int32
	FocusRegen   int32

	Agility    int32
	Evasion    int32
	Accuracy   int32
	Initiative int32

	// a list of current modifiers that affect the various stats
	Modifiers map[string]int32

	Gold int32
	XP   int32

	Script string

	ls *lua.LState

	proto *lua.FunctionProto

	haveTick       bool
	haveGetAction  bool
	haveRound      bool
	haveInitiative bool
	luaFns         map[string]lua.LValue
}

// init should be called when first loading in the script. It creates a "base copy" of the creature,
// with all the stats set up and the script read in. The lua State isn't set up yet though -- that's
// what `spawn()` is for -- it copies the creature and calls `parse()`.
func (c *creature) init() error {
	c.Statuses = []Status{}
	c.Modifiers = map[string]int32{}

	c.MaxVitality = vitalityForStat(c.Strength)
	c.CurrentVitality = c.MaxVitality

	c.CurrentFocus = focusForStat(c.Agility)
	c.MaxFocus = c.CurrentFocus
	c.Evasion = evasionForStat(c.Agility)

	c.Willpower = willForStat(c.Intelligence)
	c.Accuracy = statToMod[c.Agility]

	c.FocusRegen = statToMod[c.Intelligence]
	c.CombatDamageBase = statToMod[c.Strength]
	c.VitalityRegen = statToMod[c.Strength]
	c.Initiative = statToMod[c.Agility]

	c.luaFns = map[string]lua.LValue{}

	bits := bytes.NewBufferString(c.Script)
	chunk, err := parse.Parse(bits, c.Name)
	if err != nil {
		return err
	}
	proto, err := lua.Compile(chunk, c.Name)
	if err != nil {
		return err
	}
	c.proto = proto

	return nil
}

// spawn creates a 'live' copy of the creature, where the lua script has been initialized
func (c *creature) spawn() (*creature, error) {
	po := *c.Position
	cr := &creature{
		Id:          c.Id,
		Name:        c.Name,
		Description: c.Description,
		Position:    &po,
		Statuses:    []Status{},

		Strength:        c.Strength,
		MaxVitality:     c.MaxVitality,
		CurrentVitality: c.CurrentVitality,
		VitalityRegen:   c.VitalityRegen,

		Intelligence: c.Intelligence,
		CurrentFocus: c.CurrentFocus,
		MaxFocus:     c.MaxFocus,
		Willpower:    c.Willpower,
		FocusRegen:   c.FocusRegen,

		Agility:    c.Agility,
		Evasion:    c.Evasion,
		Accuracy:   c.Accuracy,
		Initiative: c.Initiative,

		Modifiers: map[string]int32{},
		Gold:      c.Gold,
		XP:        c.XP,
		proto:     c.proto,
		luaFns:    map[string]lua.LValue{},
	}
	err := cr.parse()
	if err != nil {
		return nil, err
	}
	return cr, nil
}

// parse uses the function proto to set up the lua state
func (c *creature) parse() error {
	//lua.NewState(opts ...lua.Options) *lua.LState
	l := lua.NewState()
	lfunc := l.NewFunctionFromProto(c.proto)
	l.Push(lfunc)
	if err := l.PCall(0, lua.MultRet, nil); err != nil {
		return err
	}
	c.ls = l
	if checkForFunction("tick", l) {
		c.haveTick = true
		c.luaFns["tick"] = l.GetGlobal("tick")
	}

	if checkForFunction("getAction", l) {
		c.haveGetAction = true
		c.luaFns["getAction"] = l.GetGlobal("getAction")
	}

	if checkForFunction("round", l) {
		c.haveRound = true
		c.luaFns["round"] = l.GetGlobal("round")
	}

	if checkForFunction("initiative", l) {
		c.haveInitiative = true
		c.luaFns["initiative"] = l.GetGlobal("initiative")
	}

	return nil
}

// checkForFunction ...
func checkForFunction(name string, state *lua.LState) bool {
	fn := state.GetGlobal(name)
	return !lua.LVIsFalse(fn) && fn.Type() == lua.LTFunction
}

// tick ...
func (c *creature) tick() error {
	if c.haveTick {
		c.ls.SetGlobal("creature", luar.New(c.ls, c))
		call := lua.P{
			Fn:      c.luaFns["tick"],
			NRet:    1,
			Protect: true,
		}

		if err := c.ls.CallByParam(call); err != nil {
			return err
		}

		ret := c.ls.Get(-1)
		c.ls.Pop(1)
		log.Printf("creature tick output: %#v", ret)
	}
	return nil
}

// getAction ...
func (c *creature) getAction(inp *endless.Input) (action, error) {
	if c.haveGetAction {
		c.ls.SetGlobal("creature", luar.New(c.ls, c))
		call := lua.P{
			Fn:      c.luaFns["getAction"],
			NRet:    1,
			Protect: true,
		}

		if err := c.ls.CallByParam(call, luar.New(c.ls, inp)); err != nil {
			return nil, err
		}

		ret := c.ls.Get(-1)
		c.ls.Pop(1)
		log.Printf("creature getAction output: %#v", ret)
	}
	return nil, nil
}

// round ...
func (c *creature) round() error {
	return nil
}

// iniative ...
func (c *creature) iniative() int {
	if c.haveInitiative {
		c.ls.SetGlobal("creature", luar.New(c.ls, c))
		call := lua.P{
			Fn:      c.luaFns["initiative"],
			NRet:    1,
			Protect: true,
		}
		if err := c.ls.CallByParam(call); err != nil {
			return 0
		}
		ret := c.ls.Get(-1)
		c.ls.Pop(1)
		if i, ok := ret.(lua.LNumber); ok {
			return int(i)
		}
	}
	return 1
}
