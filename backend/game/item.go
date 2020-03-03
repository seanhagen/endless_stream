package game

import lua "github.com/yuin/gopher-lua"

type item struct {
	script string
	ls     *lua.LState

	Name        string
	Description string
	ScriptName  string `mapstructure:"script"`
}

// itemMap is used by the game engine to hold all the possible items
type itemMap map[string]item

// inventory is used in a player struct to hold that player's inventory
// (so it has limits as to how many things it can hold)
type inventory map[string]item

// add attemps to add the given item to a players inventory
// it returns false if there wasn't enough room
func (inv inventory) add(i item) bool {
	return false
}

// remove ...
func (inv inventory) remove(i item) {}
