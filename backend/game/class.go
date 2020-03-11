package game

import (
	"fmt"

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

// createPlayerClass ...
func (cm classMap) createPlayerClass(pid string, ct endless.ClassType, g *Game) (*player, error) {
	id := cm.getIdByName(ct)
	if id == "" {
		return nil, fmt.Errorf("no class found for type '%v'", ct.String())
	}
	c, ok := cm[id]
	if !ok {
		return nil, fmt.Errorf("no class in map for id '%v'", id)
	}

	return c.getPlayer(pid, ct, g)
}

// createAI ...
func (cm classMap) createAI(pid string, ct endless.ClassType, g *Game) (*player_ai, error) {
	id := cm.getIdByName(ct)
	if id == "" {
		return nil, fmt.Errorf("no class found for type '%v'", ct.String())
	}
	c, ok := cm[id]
	if !ok {
		return nil, fmt.Errorf("no class in map for id '%v'", id)
	}

	return c.getAIPlayer(pid, ct, g)
}

// getIdByName ...
func (cm classMap) getIdByName(ct endless.ClassType) string {
	for id, c := range cm {
		if c.Name == ct.String() {
			return id
		}
	}
	return ""
}

// createCreature ...
func (c class) createCreature(pid, script string, g *Game) (*creature, error) {
	cr := &creature{
		Id:           pid,
		Name:         c.Name,
		Description:  c.Description,
		Strength:     c.Strength,
		Intelligence: c.Intelligence,
		Agility:      c.Agility,
		Position:     endless.Position_Left,
		Modifiers:    map[string]int32{},
		Script:       script,
		Level:        1,
		MType:        endless.Type_HumanPlayer,
	}
	cr.setup()
	return cr.spawn(g)
}

// getPlayer ...
func (c class) getPlayer(pid string, ct endless.ClassType, g *Game) (*player, error) {
	cr, err := c.createCreature(pid, c.baseScript, g)
	if err != nil {
		return nil, err
	}
	// cr.Skills = g.entityCollection.Skills.getClassSkills(c.Name)

	pl := player{
		creature:  cr,
		class:     ct,
		inventory: inventory{},
	}

	return &pl, nil
}

// getAIPlayer ...
func (c class) getAIPlayer(pid string, ct endless.ClassType, g *Game) (*player_ai, error) {
	cr, err := c.createCreature(pid, c.aiScript, g)
	if err != nil {
		return nil, err
	}
	pl := player_ai{
		creature: cr,
		class:    ct,
		// skills:    g.entityCollection.Skills.getClassSkills(c.Name),
		inventory: inventory{},
	}
	return &pl, nil
}
