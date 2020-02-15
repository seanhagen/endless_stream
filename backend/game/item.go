package game

import lua "github.com/yuin/gopher-lua"

type item struct {
	script string
	ls     *lua.LState

	Name        string
	Description string
	ScriptName  string `mapstructure:"script"`
}

type itemMap map[string]item
