package game

type skillConfig struct {
	Name        string
	Description map[int]string
	Type        string
	Cost        int32
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
func (sc skillsConfig) loadScripts(scriptLoad func(string) string) (skillMap, error) {
	out := skillMap{}

	for class, skills := range sc {
		cm, ok := out[class]
		if !ok {
			cm = charSkillMap{}
		}
		for id, s := range skills {
			script := scriptLoad(s.Script)

			sk := skill{skillConfig: s, script: script}
			if err := sk.init(); err != nil {
				return nil, err
			}
			cm[id] = &sk
		}
		out[class] = cm
	}
	return out, nil
}
