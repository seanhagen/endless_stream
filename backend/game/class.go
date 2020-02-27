package game

import (
	"github.com/seanhagen/endless_stream/backend/endless"
	lua "github.com/yuin/gopher-lua"
)

type class struct {
	baseScript string
	aiScript   string
	ls         *lua.LState

	Name          string
	Description   string
	Strength      int32
	Intelligence  int32
	Agility       int32
	ScriptName    string   `mapstructure:"script"`
	AIScriptName  string   `mapstructure:"ai"`
	StartingItems []string `mapstructure:"startingItems"`
}

type classMap map[string]class

// getPlayer ...
func (cm classMap) getPlayer(cid string, isAI bool, ec EntityCollection) *player {
	if c, ok := cm[cid]; ok {
		return c.getPlayer(cid, isAI, ec)
	}
	return nil
}

// getPlayer ...
func (c class) getPlayer(cid string, isAI bool, ec EntityCollection) *player {
	p := endless.Position_Left
	cr := creature{
		Id:           cid,
		Name:         c.Name,
		Description:  c.Description,
		Strength:     c.Strength,
		Intelligence: c.Intelligence,
		Agility:      c.Agility,
		Position:     &p,
		Modifiers:    map[string]int32{},
	}
	cr.init()

	cl := endless.Class{Class: endless.ClassType(endless.ClassType_value[c.Name])}
	pl := player{
		creature:  cr,
		class:     cl,
		isAI:      isAI,
		level:     1,
		skills:    ec.Skills.getClassSkills(c.Name),
		inventory: inventory{},
	}

	return &pl
}
