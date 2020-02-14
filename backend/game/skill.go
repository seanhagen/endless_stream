package game

import lua "github.com/yuin/gopher-lua"

type skill struct {
	skillConfig

	level  int
	script string
	ls     *lua.LState
}

type skillMap map[string]skill

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

// toSkillMap ...
func (sc skillsConfig) toSkillMap(scriptLoad func(string) string) skillMap {
	out := skillMap{}

	for _, skills := range sc {
		for id, s := range skills {
			script := scriptLoad(s.Script)
			out[id] = skill{skillConfig: s, script: script}
		}
	}

	return out
}
