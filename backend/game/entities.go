package game

import (
	"fmt"
	"log"
	"strings"

	"github.com/gobuffalo/packd"
	"github.com/spf13/viper"
)

type EntityCollection struct {
	Items    itemMap
	Monsters monsterMap
	Skills   skillMap
	Classes  classMap
	Statuses baseStatuses
}

const EntityDir = "entities"

func checkForScriptDirs(names []string, scripts Box) error {
	for _, n := range names {
		if !scripts.HasDir(n) {
			return fmt.Errorf("no script dir '%v'", n)
		}
	}
	return nil
}

func loadScript(scriptType, name string, scripts Box) string {
	s, err := scripts.FindString(fmt.Sprintf("%v/%v.lua", scriptType, name))
	if err != nil {
		log.Printf("unable to load script %v/%v, error: %v", scriptType, name, err)
		return ""
	}
	return s
}

// SetupEntityCollection ...
func SetupEntityCollection(scripts, entities Box) (EntityCollection, error) {
	log.Printf("Loading entities")
	out := EntityCollection{}

	scriptTypes := []string{}

	v := viper.New()
	v.SetConfigType("yaml")
	err := entities.Walk(func(n string, f packd.File) error {
		if err := v.MergeConfig(f); err != nil {
			return err
		}
		scriptTypes = append(scriptTypes, strings.Replace(n, ".yaml", "", -1))
		return nil
	})
	if err != nil {
		return out, err
	}

	err = checkForScriptDirs(scriptTypes, scripts)
	if err != nil {
		return out, err
	}

	for t := range v.AllSettings() {
		fmt.Printf("settings: %v\n", t)
		switch t {
		case "statuses":
			bs := baseStatuses{}
			err := v.UnmarshalKey(t, &bs)
			if err != nil {
				return out, err
			}
			for id, s := range bs {
				s.script = loadScript("statuses", s.ScriptName, scripts)
				bs[id] = s
			}
			out.Statuses = bs

		case "monsters":
			ms := map[string]monsterBase{}
			err := v.UnmarshalKey(t, &ms)
			if err != nil {
				return out, err
			}

			monsters := monsterMap{}
			for k, m := range ms {
				sc := loadScript("monsters", m.Script, scripts)
				monsters[k] = createMonster(k, m, sc)
			}
			out.Monsters = monsters

		case "skills":
			sl := skillsConfig{}
			err := v.UnmarshalKey(t, &sl)
			if err != nil {
				return out, err
			}
			out.Skills = sl.loadScripts(func(n string) string {
				return loadScript("skills", n, scripts)
			})

		case "items":
			im := itemMap{}
			err := v.UnmarshalKey(t, &im)
			if err != nil {
				return out, err
			}
			for k, i := range im {
				i.script = loadScript("items", i.ScriptName, scripts)
				im[k] = i
			}

			out.Items = im

		case "classes":
			cm := classMap{}
			err := v.UnmarshalKey(t, &cm)
			if err != nil {
				return out, err
			}
			for k, i := range cm {
				i.baseScript = loadScript("class", i.ScriptName, scripts)
				i.aiScript = loadScript("ai", i.AIScriptName, scripts)
				cm[k] = i
			}
			out.Classes = cm
		}
	}

	return out, nil
}
