package game

import (
	"bytes"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/seanhagen/endless_stream/backend/endless"
	lua "github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
	luar "layeh.com/gopher-luar"
)

var _ actor = &creature{}
var _ entity = &creature{}

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

	level    int32
	mType    endless.Type
	isFlying bool

	haveTick       bool
	haveGetAction  bool
	haveRound      bool
	haveInitiative bool
	haveTakeDmg    bool
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

	if checkForFunction("takeDamage", l) {
		c.haveTakeDmg = true
		c.luaFns["takeDamage"] = l.GetGlobal("takeDamage")
	}

	return nil
}

// checkForFunction ...
func checkForFunction(name string, state *lua.LState) bool {
	fn := state.GetGlobal(name)
	return !lua.LVIsFalse(fn) && fn.Type() == lua.LTFunction
}

// callFn ...
func (c *creature) callFn(name string, numRet int, args ...interface{}) ([]lua.LValue, error) {
	c.ls.SetGlobal("creature", luar.New(c.ls, *c))
	call := lua.P{
		Fn:      c.luaFns[name],
		NRet:    numRet,
		Protect: true,
	}

	out := []lua.LValue{}

	inArgs := []lua.LValue{}
	if len(args) > 0 {
		for _, v := range args {
			switch x := v.(type) {
			case string:
				inArgs = append(inArgs, lua.LString(x))
			case int:
				inArgs = append(inArgs, lua.LNumber(x))
			case float32:
				inArgs = append(inArgs, lua.LNumber(x))
			case float64:
				inArgs = append(inArgs, lua.LNumber(x))
			case int32:
				inArgs = append(inArgs, lua.LNumber(x))
			}
		}
	}

	if err := c.ls.CallByParam(call, inArgs...); err != nil {
		return out, err
	}

	for i := numRet; i > 0; i-- {
		ret := c.ls.Get(-1)
		out = append(out, ret)
		c.ls.Pop(1)
	}
	return out, nil
}

// tick ...
func (c *creature) tick() (*endless.EventMessage, error) {
	if c.haveTick {
		out, err := c.callFn("tick", 1)
		if err != nil {
			return nil, err
		}
		log.Printf("creature tick, result: %v", spew.Sdump(out))
	}
	return nil, nil
}

// round ...
func (c *creature) round() (*endless.EventMessage, error) {
	if c.haveRound {
		out, err := c.callFn("round", 1)
		if err != nil {
			return nil, err
		}
		log.Printf("creature round, result: %v", spew.Sdump(out))
	}
	return nil, nil
}

// iniative ...
func (c *creature) initiative() int {
	if c.haveInitiative {
		out, err := c.callFn("initiative", 1)
		if err != nil {
			log.Printf("unable to get initiative: %v", err)
			return 20
		}
		if len(out) <= 0 {
			return 20
		}
		if lua.LVIsFalse(out[0]) {
			return 20
		}
		return int(lua.LVAsNumber(out[0]))
	}
	return 20
}

// health ...
func (c *creature) health() (int32, int32) {
	return c.CurrentVitality, c.MaxVitality
}

// takeDamage ...
func (c *creature) takeDamage(amount, accuracy int32) *endless.EventMessage {
	if c.haveTakeDmg {
		out, err := c.callFn("takeDamage", 1, amount, accuracy)
		if err != nil {
			log.Printf("Unable to take damage using function: %v", err)
			return nil
		}
		if len(out) <= 0 {
			log.Printf("takeDamage function didn't return ammount")
			return nil
		}
		dmg := lua.LVAsNumber(out[0])
		c.CurrentVitality -= int32(dmg)
		return nil
	}
	if accuracy >= c.Evasion {
		c.CurrentVitality -= amount
	}
	return nil
}

// apply ...
func (cr *creature) apply(am actionMessage, g *Game) error {
	return nil
}

// act ...
func (cr *creature) act() actionMessage {
	return skipMsg{}
}

// // getAction ...
// func (c *creature) getAction(inp *endless.Input) (action, error) {
// 	if c.haveGetAction {
// 		c.ls.SetGlobal("creature", luar.New(c.ls, c))
// 		call := lua.P{
// 			Fn:      c.luaFns["getAction"],
// 			NRet:    1,
// 			Protect: true,
// 		}
// 		if err := c.ls.CallByParam(call, luar.New(c.ls, inp)); err != nil {
// 			return nil, err
// 		}
// 		ret := c.ls.Get(-1)
// 		c.ls.Pop(1)
// 		log.Printf("creature getAction output: %#v", ret)
// 	}
// 	return nil, nil
// }
