package game

import (
	"fmt"

	"github.com/seanhagen/endless_stream/backend/endless"
	lua "github.com/yuin/gopher-lua"
)

type skill struct {
	skillConfig

	level  int
	script string
	ls     *lua.LState
}

type charSkillMap map[string]skill

type skillMap map[string]charSkillMap

// getClassSkills ...
func (sc skillMap) getClassSkills(c string) charSkillMap {
	return sc[c]
}

// init ...
func (s *skill) init() error {
	l := lua.NewState()
	if err := l.DoString(s.script); err != nil {
		return err
	}

	if !checkForFunction("activate", l) {
		return fmt.Errorf("skill script requires function 'activate', not found in script")
	}
	return nil
}

type runSkill struct {
	tgts []string
	cst  int32
}

// targets ...
func (rs runSkill) targets() []string {
	return rs.tgts
}

// cost ...
func (rs runSkill) cost() (int32, actionType) {
	return rs.cst, action_basic
}

// apply ...
func (rs runSkill) apply(cr *creature, g *Game) error {
	return nil
}

// output ...
func (rs runSkill) output() *endless.EventMessage {
	return nil
}
