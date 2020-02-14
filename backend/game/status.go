package game

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type baseStatuses map[string]Status

type Status struct {
	script string
	ls     *lua.LState
	cr     *creature

	ScriptName  string
	Name        string
	Description string
	Alerts      []string
}

// assignStatusToCreature ...
func (bs baseStatuses) assignStatusToCreature(id string, c *creature) error {
	s, ok := bs[id]
	if !ok {
		return fmt.Errorf("no status with id '%v'", id)
	}
	return s.build(c)
}

// build ...
func (s Status) build(c *creature) error {
	x := Status{cr: c}
	err := x.loadScript(s.script)
	if err != nil {
		return err
	}
	c.Statuses = append(c.Statuses, x)
	return nil
}

// loadScript ...
func (s *Status) loadScript(in string) error {
	l := lua.NewState()
	l.SetGlobal("creature", luar.New(l, s.cr))
	if err := l.DoString(in); err != nil {
		return err
	}
	s.script = in
	s.ls = l

	init := l.GetGlobal("init")
	if !lua.LVIsFalse(init) {
		if init.Type() == lua.LTFunction {
			call := lua.P{
				Fn:      init,
				NRet:    1,
				Protect: true,
			}
			if err := l.CallByParam(call); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("init isn't a function")
		}
	}

	tick := l.GetGlobal("tick")
	if lua.LVIsFalse(tick) || tick.Type() != lua.LTFunction {
		return fmt.Errorf("no tick function found")
	}

	return nil
}

// tick ...
func (s *Status) tick() (bool, error) {
	if err := s.ls.CallByParam(lua.P{
		Fn:      s.ls.GetGlobal("tick"),
		NRet:    1,
		Protect: true,
	}); err != nil {
		return false, err
	}

	ret := s.ls.Get(-1)
	s.ls.Pop(1)

	if ret.Type() != lua.LTBool {
		return false, fmt.Errorf("tick didn't return boolean")
	}
	return lua.LVAsBool(ret), nil
}
