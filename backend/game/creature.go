package game

import (
	"github.com/seanhagen/endless_stream/backend/endless"
	lua "github.com/yuin/gopher-lua"
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
}

// init ...
func (c *creature) init() {
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
}
