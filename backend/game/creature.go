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

	Skills charSkillMap

	ls *lua.LState

	proto *lua.FunctionProto

	Level    int32
	MType    endless.Type
	IsFlying bool

	haveTick       bool
	haveGetAction  bool
	haveRound      bool
	haveInitiative bool
	haveTakeDmg    bool
	luaFns         map[string]lua.LValue
}

// setup should be called when first loading in the script. It creates a "base copy" of the creature,
// with all the stats set up and the script read in. The lua State isn't set up yet though -- that's
// what `spawn()` is for -- it copies the creature and calls `parse()`.
func (cr *creature) setup() error {
	cr.Statuses = []Status{}
	cr.Modifiers = map[string]int32{}

	cr.MaxVitality = vitalityForStat(cr.Strength)
	cr.CurrentVitality = cr.MaxVitality

	cr.CurrentFocus = focusForStat(cr.Agility)
	cr.MaxFocus = cr.CurrentFocus
	cr.Evasion = evasionForStat(cr.Agility)

	cr.Willpower = willForStat(cr.Intelligence)
	cr.Accuracy = statToMod[cr.Agility]

	cr.FocusRegen = statToMod[cr.Intelligence]
	cr.CombatDamageBase = statToMod[cr.Strength]
	cr.VitalityRegen = statToMod[cr.Strength]
	cr.Initiative = statToMod[cr.Agility]

	cr.luaFns = map[string]lua.LValue{}

	bits := bytes.NewBufferString(cr.Script)
	chunk, err := parse.Parse(bits, cr.Name)
	if err != nil {
		return err
	}
	proto, err := lua.Compile(chunk, cr.Name)
	if err != nil {
		return err
	}
	cr.proto = proto
	return nil
}

// spawn creates a 'live' copy of the creature, where the lua script has been initialized
func (cr *creature) spawn(g *Game) (*creature, error) {
	skills, err := g.entityCollection.Skills.getClassSkills(cr.Name, g)
	if err != nil {
		return nil, err
	}

	po := *cr.Position
	crn := &creature{
		Id:          cr.Id,
		Name:        cr.Name,
		Description: cr.Description,
		Position:    &po,
		Statuses:    []Status{},

		Strength:        cr.Strength,
		MaxVitality:     cr.MaxVitality,
		CurrentVitality: cr.CurrentVitality,
		VitalityRegen:   cr.VitalityRegen,

		Intelligence: cr.Intelligence,
		CurrentFocus: cr.CurrentFocus,
		MaxFocus:     cr.MaxFocus,
		Willpower:    cr.Willpower,
		FocusRegen:   cr.FocusRegen,

		Agility:    cr.Agility,
		Evasion:    cr.Evasion,
		Accuracy:   cr.Accuracy,
		Initiative: cr.Initiative,

		Modifiers: map[string]int32{},
		Gold:      cr.Gold,
		XP:        cr.XP,
		proto:     cr.proto,
		luaFns:    map[string]lua.LValue{},

		Skills: skills,
	}
	err = crn.parse(g)
	if err != nil {
		return nil, err
	}
	return crn, nil
}

// parse uses the function proto to set up the lua state
func (cr *creature) parse(g *Game) error {
	//lua.NewState(opts ...lua.Options) *lua.LState
	l := lua.NewState()
	g.setupFunctions(l)

	lfunc := l.NewFunctionFromProto(cr.proto)
	l.Push(lfunc)
	if err := l.PCall(0, lua.MultRet, nil); err != nil {
		return err
	}
	cr.ls = l
	if checkForFunction("tick", l) {
		cr.haveTick = true
		cr.luaFns["tick"] = l.GetGlobal("tick")
	}

	if checkForFunction("getAction", l) {
		cr.haveGetAction = true
		cr.luaFns["getAction"] = l.GetGlobal("getAction")
	}

	if checkForFunction("round", l) {
		cr.haveRound = true
		cr.luaFns["round"] = l.GetGlobal("round")
	}

	if checkForFunction("initiative", l) {
		cr.haveInitiative = true
		cr.luaFns["initiative"] = l.GetGlobal("initiative")
	}

	if checkForFunction("takeDamage", l) {
		cr.haveTakeDmg = true
		cr.luaFns["takeDamage"] = l.GetGlobal("takeDamage")
	}

	return nil
}

// checkForFunction ...
func checkForFunction(name string, state *lua.LState) bool {
	fn := state.GetGlobal(name)
	return !lua.LVIsFalse(fn) && fn.Type() == lua.LTFunction
}

// callFn ...
func (cr *creature) callFn(name string, numRet int, args ...interface{}) ([]interface{}, error) {
	cr.ls.SetGlobal("creature", luar.New(cr.ls, *cr))
	call := lua.P{
		Fn:      cr.luaFns[name],
		NRet:    numRet,
		Protect: true,
	}

	out := []interface{}{}

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
			default:
				inArgs = append(inArgs, luar.New(cr.ls, x))
			}
		}
	}

	if err := cr.ls.CallByParam(call, inArgs...); err != nil {
		return out, err
	}

	for i := 0; i < numRet; i++ {
		ret := cr.ls.Get(-1)
		switch ret.Type().String() {
		case "table":
			t := cr.ls.CheckTable(-1)
			// spew.Dump(t)
			out = append(out, t)
		default:
			out = append(out, ret)
		}

		// fmt.Printf("return value %v is %v\n", i, ret.Type().String())
		cr.ls.Pop(1)
	}
	return out, nil
}

// apply ...
func (cr *creature) apply(from *creature, am actionMessage, g *Game) error {
	return am.apply(from, cr, g)
}

// act is here so that it'll catch if an monster or player doesn't implement this method
func (cr *creature) act(ws *waveState) actionMessage {
	if cr.haveGetAction {
		ws.register(cr)

		// getAction should return: id of skill to use, array of target ids
		out, err := cr.callFn("getAction", 2)
		if err != nil {
			log.Printf("unable to get action from script: %v", err)
			return skipMsg{}
		}

		if len(out) != 2 {
			log.Printf("should have 2 results from getAction, have %v", len(out))
			return skipMsg{}
		}

		st := out[1].(lua.LValue)
		skillId := lua.LVAsString(st)

		if skillId == "skip" {
			return skipMsg{}
		}

		targets := []string{}
		tbl := out[0].(*lua.LTable)
		tbl.ForEach(func(_, v lua.LValue) {
			targets = append(targets, lua.LVAsString(v))
			// fmt.Printf("table forEach: \n\ta: %v\tb: %v\n\n", spew.Sdump(a), spew.Sdump(b))
		})

		sk, ok := cr.Skills[skillId]
		if !ok {
			log.Printf("creature has no skill %v", skillId)
			return skipMsg{}
		}

		if sk.Level <= 0 {
			log.Printf("skill has no level")
			return skipMsg{}
		}

		return sk.getRunSkill(targets)
	}
	return skipMsg{}
}

// tick ...
func (cr *creature) tick() (*endless.EventMessage, error) {
	if cr.haveTick {
		out, err := cr.callFn("tick", 1)
		if err != nil {
			return nil, err
		}
		log.Printf("creature tick, result: %v", spew.Sdump(out))
	}
	return nil, nil
}

// round ...
func (cr *creature) round() (*endless.EventMessage, error) {
	// apply any status effects ( remove any that are finished )
	// if the creature has died, return a message
	// if it hasn't died:
	if cr.haveRound {
		// it has a round script, run it
		out, err := cr.callFn("round", 1)
		if err != nil {
			return nil, err
		}
		log.Printf("creature round, result: %v", spew.Sdump(out))
	}
	return nil, nil
}

// ID ...
func (cr *creature) ID() string {
	return cr.Id
}

// Health ...
func (cr *creature) Health() (int32, int32) {
	return cr.CurrentVitality, cr.MaxVitality
}

// Type ...
func (cr *creature) Type() endless.Type {
	return cr.MType
}

// iniative ...
func (cr *creature) initiative() int {
	if cr.haveInitiative {
		out, err := cr.callFn("initiative", 1)
		if err != nil {
			log.Printf("unable to get initiative: %v", err)
			return 20
		}
		if len(out) <= 0 {
			return 20
		}
		v := out[0].(lua.LValue)
		if lua.LVIsFalse(v) {
			return 20
		}
		return int(lua.LVAsNumber(v))
	}
	return 20
}

// takeDamage ...
func (cr *creature) takeDamage(amount, accuracy int32) *endless.EventMessage {
	if cr.haveTakeDmg {
		out, err := cr.callFn("takeDamage", 1, amount, accuracy)
		if err != nil {
			log.Printf("Unable to take damage using function: %v", err)
			return nil
		}
		if len(out) <= 0 {
			log.Printf("takeDamage function didn't return ammount")
			return nil
		}
		v := out[0].(lua.LValue)
		dmg := lua.LVAsNumber(v)
		cr.CurrentVitality -= int32(dmg)
		return nil
	}
	if accuracy >= cr.Evasion {
		cr.CurrentVitality -= amount
	}
	return nil
}
