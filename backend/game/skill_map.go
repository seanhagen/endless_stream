package game

import (
	"bytes"
	"fmt"

	"github.com/seanhagen/endless_stream/backend/endless"
	lua "github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
)

var _ actionMessage = &runSkill{}

type skill struct {
	skillConfig

	Level  int
	script string

	proto *lua.FunctionProto
	ls    *lua.LState
}

type charSkillMap map[string]*skill

type skillMap map[string]charSkillMap

// getClassSkills ...
func (sc skillMap) getClassSkills(c string, g *Game) (charSkillMap, error) {
	out := charSkillMap{} // map[string]*skill

	for id, sk := range sc[c] {
		ns, err := sk.spawn(g)
		if err != nil {
			return nil, err
		}

		out[id] = ns
	}
	return out, nil
}

// init ...
func (s *skill) init() error {
	bits := bytes.NewBufferString(s.script)
	chunk, err := parse.Parse(bits, fmt.Sprintf("skill-%v", s.skillConfig.Name))
	if err != nil {
		return err
	}

	proto, err := lua.Compile(chunk, s.skillConfig.Name)
	if err != nil {
		return err
	}

	s.proto = proto
	return nil
}

// spawn ...
func (s *skill) spawn(g *Game) (*skill, error) {
	l := lua.NewState()
	g.setupFunctions(l)
	lfunc := l.NewFunctionFromProto(s.proto)
	l.Push(lfunc)
	if err := l.PCall(0, lua.MultRet, nil); err != nil {
		return nil, err
	}

	if !checkForFunction("activate", l) {
		return nil, fmt.Errorf("skill script requires function 'activate', not found in script")
	}

	ns := skill{
		skillConfig: s.skillConfig,
		Level:       s.Level,
		script:      s.script,
		proto:       s.proto,
		ls:          l,
	}

	return &ns, nil
}

// // activate ...
// func (s *skill) activate() error {
// 	return nil
// }

// cost ...
func (s *skill) cost() (int32, actionType) {
	return s.Cost, action_basic
}

// apply ...
func (s *skill) apply(from, to *creature, g *Game) error {
	// from/to.Statuses -- list of statuses affecting the creature
	// from/to.Modifiers map[string]int32 modifiers for various attributes
	return nil
}

// output ...
func (s *skill) output() *endless.EventMessage {
	return nil
}

// getRunSkill ...
func (s *skill) getRunSkill(targets []string) runSkill {
	return runSkill{s, targets}
}

type runSkill struct {
	*skill
	tgts []string
}

// targets ...
func (rs runSkill) targets() []string {
	return rs.tgts
}
