package game

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/seanhagen/endless_stream/backend/endless"
)

func TestWaveStateGetCurrent(t *testing.T) {
	ws := newWaveState()

	g := &Game{}

	bsks := `function activate(from, to) end`
	sk := makeTestSkill(t, g, bsks, 1)

	g.entityCollection = EntityCollection{
		Skills: skillMap{
			"Rat": charSkillMap{"attack": sk},
		},
	}

	expIn := 5
	ratAI := fmt.Sprintf("function initiative() return %v end", expIn)

	rid := "a26fd574-e65b-432c-9611-4e8299ea9e3c"
	rat := makeTestMonster(t, g, "Rat", rid, ratAI, endless.Position_Right)
	if err := ws.addActor(rat); err != nil {
		t.Fatalf("unable to add rat to wave state: %v", err)
	}

	rid2 := "e9a0e49f-ff92-4f5e-b824-07f31d2b96f3"
	rat2 := makeTestMonster(t, g, "Rat", rid2, ratAI, endless.Position_Right)
	if err := ws.addActor(rat2); err != nil {
		t.Fatalf("unable to add second rat to wave state: %v", err)
	}

	if err := ws.waveStart(); err != nil {
		t.Fatalf("uanble to start wave: %v", err)
	}

	c := ws.current()
	if c == nil {
		t.Fatalf("expected actor got nil")
	}

	if cid := c.ID(); cid != rid {
		t.Errorf("wrong id, expected %v got %v", rid, cid)
	}

	if ws.current_initiative != expIn {
		t.Errorf("expected to be on step %v, actually on step: %v", expIn, ws.current_initiative)
	}
}

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

			err = ws.act()
			if err != nil {
				t.Fatalf("unable to act: %v", err)
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

func TestWaveStateCurrentNil(t *testing.T) {
	ws := newWaveState()
	if a := ws.current(); a != nil {
		t.Errorf("expected nil, got an actor?")
	}
}

func TestWaveStateGetPlayers(t *testing.T) {
	ws := newWaveState()

	g := &Game{}

	pid1 := "7451595d-6260-4e9f-823a-23aaf10f1038"
	pid2 := "818cdec1-b4d3-42fc-a135-370899afff15"

	bsks := `function activate(from, to) end`
	sk := makeTestSkill(t, g, bsks, 1)

	g.entityCollection = EntityCollection{
		Skills: skillMap{
			"Fighter": charSkillMap{"attack": sk},
			"Cleric":  charSkillMap{"attack": sk},
			"Rat":     charSkillMap{"attack": sk},
		},
		Classes: classMap{
			"de8d9f47-2a1c-4972-a7ae-742433613a3d": class{Name: "Fighter"},
			"877e0dd1-554e-49f2-8fdb-f867f490ba46": class{Name: "Cleric"},
		},
	}

	ratAI := `function initiative() return 1 end
function getAction()
pids = getPlayers()

targets = {}
idx = 0
for i,l in pairs(pids) do
  targets[idx] = l
  idx = idx+1
end

return "attack", targets
end`

	rid := "a26fd574-e65b-432c-9611-4e8299ea9e3c"
	rat := makeTestMonster(t, g, "Rat", rid, ratAI, endless.Position_Right)
	if err := ws.addActor(rat); err != nil {
		t.Fatalf("unable to add rat to wave state: %v", err)
	}

	p1 := makeTestPlayer(t, g, endless.ClassType_Fighter, pid1, endless.Position_Left)
	if err := ws.addActor(p1); err != nil {
		t.Fatalf("unable to add player 1 to wave state: %v", err)
	}

	p2 := makeTestPlayer(t, g, endless.ClassType_Fighter, pid2, endless.Position_Left)
	if err := ws.addActor(p2); err != nil {
		t.Fatalf("unable to add player 2 to wave state: %v", err)
	}

	if err := ws.waveStart(); err != nil {
		t.Fatalf("unable to start wave: %v", err)
	}

	ca := ws.current()
	if ca == nil || ca.ID() != rid {
		t.Fatalf("current actor is not rat")
	}

	am := ca.act(ws)
	tgts := am.targets()
	expt := []string{pid1, pid2}

	if !stringSliceEq(tgts, expt) {
		t.Errorf("wrong targets, expected %#v got %#v", expt, tgts)
	}

}

func TestWaveStateMove(t *testing.T) {
	tests := []struct {
		start endless.Position
		end   endless.Position
		skill string
	}{
		{endless.Position_Left, endless.Position_Middle, "moveRight"},
		{endless.Position_Middle, endless.Position_Right, "moveRight"},
		{endless.Position_Right, endless.Position_Right, "moveRight"},
		{endless.Position_Right, endless.Position_Middle, "moveLeft"},
		{endless.Position_Middle, endless.Position_Left, "moveLeft"},
		{endless.Position_Left, endless.Position_Left, "moveLeft"},
	}

	moveTmpl := `function getAction()
  return "%v", {}
end`

	for i := range tests {
		tt := tests[i]
		t.Run(fmt.Sprintf("move test start %v end %v action %v", tt.start, tt.end, tt.skill), func(t *testing.T) {
			ws := newWaveState()

			g := &Game{
				entityCollection: EntityCollection{
					Skills: skillMap{
						"Rat": charSkillMap{},
					},
				},
			}

			ratAI := fmt.Sprintf(moveTmpl, tt.skill)
			rat := makeTestMonster(t, g, "Rat", "0c1f01dc-e21d-44ea-9d14-6f65c7c5ca46", ratAI, tt.start)
			if err := ws.addActor(rat); err != nil {
				t.Fatalf("unable to add rat to wave state: %v", err)
			}

			if err := ws.waveStart(); err != nil {
				t.Fatalf("unable to trigger wave start: %v", err)
			}

			if err := ws.act(); err != nil {
				t.Fatalf("unable to act: %v", err)
			}

			if ws.proceed() {
				if err := ws.process(g); err != nil {
					t.Fatalf("unable to process wave state: %v", err)
				}
			}

			at := rat.creature.Position
			if at != tt.end {
				t.Errorf("wrong position, expected %v got %v", tt.end, at)
			}
		})
	}
}

func TestWaveComplete(t *testing.T) {
	tests := []struct {
		playerHealth  int32
		monsterHealth int32
		complete      bool
	}{
		{10, 5, false},
		{10, 0, true},
		{0, 0, false},
		{0, 5, false},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(fmt.Sprintf("test p%v m%v g%v", tt.playerHealth, tt.monsterHealth, tt.complete), func(t *testing.T) {
			ws := newWaveState()
			g := &Game{
				entityCollection: EntityCollection{
					Skills: skillMap{
						"Rat":     charSkillMap{},
						"Fighter": charSkillMap{},
						"Cleric":  charSkillMap{},
						"Ranger":  charSkillMap{},
						"Wizard":  charSkillMap{},
					},
					Classes: classMap{
						"1": class{Name: "Fighter"},
						"2": class{Name: "Cleric"},
						"3": class{Name: "Ranger"},
						"4": class{Name: "Wizard"},
					},
				},
			}

			rat := makeTestMonster(t, g, "Rat", "rat1", "", endless.Position_Right)
			if err := ws.addActor(rat); err != nil {
				t.Fatalf("unable to add rat to wave state: %v", err)
			}

			fighter := makeTestPlayer(t, g, endless.ClassType_Fighter, "fighter1", endless.Position_Left)
			if err := ws.addActor(fighter); err != nil {
				t.Fatalf("unable to add fighter player to wave state: %v", err)
			}

			cleric := makeTestPlayer(t, g, endless.ClassType_Cleric, "cleric1", endless.Position_Left)
			if err := ws.addActor(cleric); err != nil {
				t.Fatalf("unable to add cleric player to wave state: %v", err)
			}

			ranger := makeTestPlayer(t, g, endless.ClassType_Ranger, "ranger1", endless.Position_Left)
			if err := ws.addActor(ranger); err != nil {
				t.Fatalf("unable to add ranger player to wave state: %v", err)
			}

			wizard := makeTestPlayer(t, g, endless.ClassType_Wizard, "wizard1", endless.Position_Left)
			if err := ws.addActor(wizard); err != nil {
				t.Fatalf("unable to add wizard player to wave state: %v", err)
			}

			fighter.creature.CurrentVitality = tt.playerHealth
			cleric.creature.CurrentVitality = tt.playerHealth
			ranger.creature.CurrentVitality = tt.playerHealth
			wizard.creature.CurrentVitality = tt.playerHealth

			rat.creature.CurrentVitality = tt.monsterHealth

			if c := ws.waveComplete(); c != tt.complete {
				t.Errorf("wrong value for complete, expected '%v' got '%v'", tt.complete, c)
			}
		})
	}
}

func TestWaveFailed(t *testing.T) {
	tests := []struct {
		playerHealth  int32
		monsterHealth int32
		complete      bool
	}{
		{10, 5, false},
		{10, 0, false},
		{0, 0, true},
		{0, 5, true},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(fmt.Sprintf("test p%v m%v g%v", tt.playerHealth, tt.monsterHealth, tt.complete), func(t *testing.T) {
			ws := newWaveState()
			g := &Game{
				entityCollection: EntityCollection{
					Skills: skillMap{
						"Rat":     charSkillMap{},
						"Fighter": charSkillMap{},
						"Cleric":  charSkillMap{},
						"Ranger":  charSkillMap{},
						"Wizard":  charSkillMap{},
					},
					Classes: classMap{
						"1": class{Name: "Fighter"},
						"2": class{Name: "Cleric"},
						"3": class{Name: "Ranger"},
						"4": class{Name: "Wizard"},
					},
				},
			}

			rat := makeTestMonster(t, g, "Rat", "rat1", "", endless.Position_Right)
			if err := ws.addActor(rat); err != nil {
				t.Fatalf("unable to add rat to wave state: %v", err)
			}

			fighter := makeTestPlayer(t, g, endless.ClassType_Fighter, "fighter1", endless.Position_Left)
			if err := ws.addActor(fighter); err != nil {
				t.Fatalf("unable to add fighter player to wave state: %v", err)
			}

			cleric := makeTestPlayer(t, g, endless.ClassType_Cleric, "cleric1", endless.Position_Left)
			if err := ws.addActor(cleric); err != nil {
				t.Fatalf("unable to add cleric player to wave state: %v", err)
			}

			ranger := makeTestPlayer(t, g, endless.ClassType_Ranger, "ranger1", endless.Position_Left)
			if err := ws.addActor(ranger); err != nil {
				t.Fatalf("unable to add ranger player to wave state: %v", err)
			}

			wizard := makeTestPlayer(t, g, endless.ClassType_Wizard, "wizard1", endless.Position_Left)
			if err := ws.addActor(wizard); err != nil {
				t.Fatalf("unable to add wizard player to wave state: %v", err)
			}

			fighter.creature.CurrentVitality = tt.playerHealth
			cleric.creature.CurrentVitality = tt.playerHealth
			ranger.creature.CurrentVitality = tt.playerHealth
			wizard.creature.CurrentVitality = tt.playerHealth

			rat.creature.CurrentVitality = tt.monsterHealth

			if c := ws.waveFailed(); c != tt.complete {
				t.Errorf("wrong value for failed, expected '%v' got '%v'", tt.complete, c)
			}
		})
	}
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
				baseScript: `function initiative() return 5 end
function getAction()
  print("[SCRIPT] fighter ID: ", creature.Id)
  print("[SCRIPT] fighter current position: ", creature.Position)

  pids = getMonsters()
  pl = entityByKey(pids[1])

  print("[SCRIPT] figher checks rat position: ", pl.Position)

  if pl.Position > creature.Position then
    print("[SCRIPT] fighter moves")
    return "moveRight", {}
  end

  print("[SCRIPT] fighter skips")
  return "skip", {}
end
`,
			},
		},
	}

	ratAI := `function initiative() return 10 end
function getAction()
  print("[SCRIPT] rat ID: ", creature.Id)
  pids = getPlayers()
  pl = entityByKey(pids[1])

  print("\n[SCRIPT] rat position: ", creature.Position, "fighter position: ", pl.Position)
  if pl.Position < creature.Position then
    print("[SCRIPT] rat moves")
    return "moveLeft", {}
  end

  print("[SCRIPT] rat skips", pids[1], pl.Position, creature.Position)
  return "skip", {}
end
`
	rat := makeTestMonster(t, g, "Rat", "0c1f01dc-e21d-44ea-9d14-6f65c7c5ca46", ratAI, endless.Position_Right)
	if err := ws.addActor(rat); err != nil {
		t.Fatalf("unable to add rat to wave state: %v", err)
	}

	fighter := makeTestPlayer(t, g, endless.ClassType_Fighter, "4230a173-d0bc-4b64-b065-5810ebe2d928", endless.Position_Left)
	if err := ws.addActor(fighter); err != nil {
		t.Fatalf("unable to add fighter player to wave state: %v", err)
	}

	fighter.creature.CurrentVitality = 10
	rat.creature.CurrentVitality = 5

	if err := ws.waveStart(); err != nil {
		t.Fatalf("unable to trigger wave start: %v", err)
	}

	maxTurns := 20
	turn := 1

	for i := turn; i < maxTurns; i++ {
		if ws.current_round == 2 {
			t.Fatalf("one round done")
		}

		if err := ws.act(); err != nil {
			t.Fatalf("unable to act: %v", err)
		}
		if ws.proceed() {

			// ca := ws.current()
			// fmt.Printf("current actor: %v\n", ca.ID())

			if err := ws.process(g); err != nil {
				t.Fatalf("unable to process wave state: %v", err)
			}

			if ws.waveComplete() || ws.waveFailed() {
				t.Errorf("wave should not be complete or failed yet")
			}

			fmt.Printf("fighter health: %v\n", fighter.creature.CurrentVitality)
			fmt.Printf("rat health: %v\n", rat.creature.CurrentVitality)

		}
		fmt.Printf("[TEST] current_round: %v\n", ws.current_round)
	}

	t.Errorf("not yet")
}
