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
						"Monster": charSkillMap{
							"07f0def4-4fe3-426a-a500-4013f89506ab": &sk,
						},
						"Monster 2": charSkillMap{
							"8d30a3cd-ec43-4142-b436-a4d245e80801": &sk,
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
}

func TestWaveStateGetPlayers(t *testing.T) {
	t.Errorf("not yet")
}

func TestWaveStateSimpleBattle(t *testing.T) {
	ws := newWaveState()

	g := &Game{}

	ratAtk := `function activate(from, to)
to.CurrentVitality = to.CurrentVitality - 1
end`

	fightAtk := `function activate(from, to)
to.CurrentVitality = to.CurrentVitality - 1
end`

	ratAtkSk := makeTestSkill(t, g, ratAtk, 1)
	fightAtkSk := makeTestSkill(t, g, fightAtk, 1)

	g.entityCollection = EntityCollection{
		Skills: skillMap{
			"Rat":     charSkillMap{"attack": ratAtkSk},
			"Fighter": charSkillMap{"attack": fightAtkSk},
		},
		Classes: classMap{
			"40f9a099-58e3-4769-8a1f-80f8bf0982fb": class{
				Name: "Fighter",
			},
		},
	}

	ratAI := `function initiative() return 10 end
function getAction()
  return "skip", {}
end
`
	rat := makeTestMonster(t, g, "Rat", "0c1f01dc-e21d-44ea-9d14-6f65c7c5ca46", ratAI, endless.Position_Right)

	if err := ws.addActor(rat); err != nil {
		t.Fatalf("unable to add rat to wave state: %v", err)
	}

	fightAI := `function initiative() return 5 end
function getAction()
  print("current position: ", creature.Position)
  return "skip", {}
end
`
	fighter := makeTestPlayer(t, g, endless.ClassType_Fighter, "4230a173-d0bc-4b64-b065-5810ebe2d928", fightAI, endless.Position_Left)

	fighter.creature.CurrentVitality = 10
	rat.creature.CurrentVitality = 5

	if err := ws.addActor(fighter); err != nil {
		t.Fatalf("unable to add fighter player to wave state: %v", err)
	}

	if err := ws.waveStart(); err != nil {
		t.Fatalf("unable to trigger wave start: %v", err)
	}

	if err := ws.tick(); err != nil {
		t.Fatalf("unable to tick: %v", err)
	}

	if ws.proceed() {
		if err := ws.process(); err != nil {
			t.Fatalf("unable to process wave state: %v", err)
		}

		if ws.waveComplete() || ws.waveFailed() {
			t.Errorf("wave should not be complete or failed yet")
		}

	}

}
