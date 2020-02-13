package game

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type status struct {
	script string
	ls     *lua.LState

	// stepInput string
	// turnDone  bool

	cr   *creature
	done bool
}

// newStatus ...
func newStatus(s string, c *creature) (*status, error) {
	st := &status{
		cr: c,
	}
	err := st.loadScript(s)
	if err != nil {
		return nil, err
	}
	return st, nil
}

// loadScript ...
func (s *status) loadScript(in string) error {
	l := lua.NewState()
	l.SetGlobal("creature", luar.New(l, s.cr))
	if err := l.DoString(in); err != nil {
		return err
	}
	s.script = in
	s.ls = l

	call := lua.P{
		Fn:      l.GetGlobal("init"),
		NRet:    1,
		Protect: true,
	}
	if err := l.CallByParam(call); err != nil {
		return err
	}

	return nil
}

// tick ...
func (s *status) tick() (bool, error) {
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
