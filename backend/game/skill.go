package game

import lua "github.com/yuin/gopher-lua"

type skill struct {
	skillConfig

	level  int
	script string
	ls     *lua.LState
}

type charSkillMap map[string]skill

type skillMap map[string]charSkillMap

type skillConfig struct {
	Name        string
	Description map[int]string
	Type        string
	Cost        int
	Script      string
}

type charSkillConfig map[string]skillConfig

type skillsConfig map[string]charSkillConfig

// toClassSkills ...
func (sc skillsConfig) toClassSkills() map[string][]string {
	out := map[string][]string{}

	for cl, skills := range sc {
		for id := range skills {
			c, ok := out[cl]
			if !ok {
				c = []string{}
			}
			c = append(c, id)
			out[cl] = c
		}
	}

	return out
}

// loadScripts ...
func (sc skillsConfig) loadScripts(scriptLoad func(string) string) skillMap {
	out := skillMap{}

	for class, skills := range sc {
		cm, ok := out[class]
		if !ok {
			cm = charSkillMap{}
		}
		for id, s := range skills {
			script := scriptLoad(s.Script)
			cm[id] = skill{skillConfig: s, script: script}
		}
		out[class] = cm
	}
	return out
}

// getClassSkills ...
func (sc skillMap) getClassSkills(c string) charSkillMap {
	return sc[c]
}
