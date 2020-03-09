package game

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/seanhagen/endless_stream/backend/endless"
)

func TestWaveStateTick(t *testing.T) {
	basicSkillScript := `function activate() end`

	tests := []struct {
		skillScript   string
		monsterScript string
		costA         int32
		costT         actionType
		ret           actionMessage
		targets       []string
	}{
		{basicSkillScript, `noMonsterAction = 1`, 0, action_basic, skipMsg{}, []string{}},
		{basicSkillScript, `function getAction() return "07f0def4-4fe3-426a-a500-4013f89506ab", {"monster-2-1"} end`, 1, action_basic, runSkill{}, []string{"monster-2-1"}},
	}

	for i, x := range tests {
		tt := x
		j := i
		t.Run(fmt.Sprintf("test %v", j), func(t *testing.T) {
			sk := skill{
				skillConfig: skillConfig{
					Name: "Basic Attack",
					Cost: 1,
				},
				Level:  1,
				script: tt.skillScript,
			}
			err := sk.init()
			if err != nil {
				t.Fatalf("unable to init skill: %v", err)
			}

			g := &Game{
				entityCollection: EntityCollection{
					Skills: skillMap{
						"Rat": charSkillMap{
							"07f0def4-4fe3-426a-a500-4013f89506ab": &sk,
						},
					},
				},
			}

			ws := newWaveState()
			mid := fmt.Sprintf("monster-1-%v", j)
			mid2 := fmt.Sprintf("monster-2-%v", j)

			m := makeTestMonster(t, g, "Monster", mid, tt.monsterScript, endless.Position_Right)
			m2 := makeTestMonster(t, g, "Monster 2", mid2, "", endless.Position_Right)

			ci := ws.current_initiative
			ws.initiative[ci] = []actor{m, m2}
			ws.Entities[mid] = m
			ws.Entities[mid2] = m2

			err = ws.tick()
			if err != nil {
				t.Fatalf("unable to tick: %v", err)
			}

			if ws.currentAction == nil && tt.ret != nil {
				t.Fatalf("current action is nil")
			}

			et := reflect.TypeOf(tt.ret)
			tp := reflect.TypeOf(ws.currentAction)
			if et.String() != tp.String() {
				t.Errorf("wrong action message type, expected '%v', got '%v'", et.String(), tp.String())
			}

			ca, ct := ws.currentAction.cost()
			if ca != tt.costA {
				t.Errorf("wrong action cost, expected '%v', got '%v'", tt.costA, ca)
			}
			if ct != tt.costT {
				t.Errorf("wrong action cost type, expected '%v', got '%v'", tt.costT, ct)
			}

			if tgts := ws.currentAction.targets(); !stringSliceEq(tgts, tt.targets) {
				t.Errorf("wrong targets, expected %#v, got %#v", tgts, tt.targets)
			}

			if !ws.proceed() {
				t.Errorf("expected true from `proceed()`, got false")
			}

		})
	}

	// m1 := makeCreature(t, g, mid, script, p)
	// m2 := makeCreature(t, g, mid2, script, p)
	// ws := newWaveState()
	// ws.Entities[mid] = m1
	// ws.Entities[mid2] = m2
	//t.Errorf("not yet")
}
