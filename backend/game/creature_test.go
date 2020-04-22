package game

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/seanhagen/endless_stream/backend/endless"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func TestCreatureInit(t *testing.T) {
	tests := []struct {
		name string
		str  int32
		int  int32
		agi  int32

		vit  int32
		foc  int32
		will int32
		eva  int32
		acc  int32
	}{
		{"rat", 0, 0, 1, 5, 10, 5, 7, -2},
		{"corvid shaman", 2, 4, 3, 15, 20, 13, 11, 0},
		{"wendigo", 6, 5, 7, 35, 40, 15, 19, 4},
	}

	for _, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", tt.name), func(t *testing.T) {
			cr := creature{
				Strength:     tt.str,
				Intelligence: tt.int,
				Agility:      tt.agi,
			}

			cr.setup()
			if cr.MaxVitality != tt.vit {
				t.Errorf("Wrong vitality, expected %v got %v", tt.vit, cr.MaxVitality)
			}

			if cr.MaxFocus != tt.foc {
				t.Errorf("Wrong focus, expected %v got %v", tt.foc, cr.MaxFocus)
			}

			if cr.Willpower != tt.will {
				t.Errorf("Wrong willpower, expected %v got %v", tt.will, cr.Willpower)
			}

			if cr.Evasion != tt.eva {
				t.Errorf("Wrong evasion, expected %v got %v", tt.eva, cr.Evasion)
			}

			if cr.Accuracy != tt.acc {
				t.Errorf("Wrong accuracy, expected %v got %v", tt.acc, cr.Accuracy)
			}

		})
	}
}

func TestCreatureInitiative(t *testing.T) {
	tests := []struct {
		name       string
		script     string
		expectInit int
	}{
		{"no initiative", "initBase = 0", 20},
		{"simple intiative", `function initiative() return 3 end`, 3},
		{"complex initiative", `function initiative() return 11 + creature.Initiative end`, 8},
		{"bad initiative script", `function initiative() end`, 20},
	}

	g := &Game{
		entityCollection: EntityCollection{
			Skills: skillMap{
				"": charSkillMap{},
			},
		},
	}

	for _, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", tt.name), func(t *testing.T) {
			b := &creature{Script: tt.script, Position: endless.Position_Right}
			err := b.setup()
			if err != nil {
				t.Fatalf("unable to initialize creature: %v", err)
			}

			cr, err := b.spawn(g)
			if err != nil {
				t.Fatalf("unable to spawn active creature: %v", err)
			}

			init := cr.initiative()
			if init != tt.expectInit {
				t.Errorf("invalid value, expected '%v' got '%v'", tt.expectInit, init)
			}
		})
	}
}

func TestCreatureTakeDamage(t *testing.T) {
	simpleScript := `function takeDamage(amt, acc)
if acc > creature.Evasion then
  return amt
end
return 0
end`

	armorReduce := `function takeDamage(amt, acc)
if acc > creature.Evasion then
  return amt - 2
end
return 0
end`

	tests := []struct {
		name         string
		script       string
		evasion      int32
		health       int32
		accuracy     int32
		amount       int32
		expectHealth int32
	}{
		{"simple", "noDamageFn = 1", 5, 10, 6, 5, 5},
		{"simple script", simpleScript, 5, 10, 6, 5, 5},
		{"armor script", armorReduce, 5, 10, 6, 5, 7},
		{"armor no dmg", armorReduce, 5, 10, 6, 2, 10},
		{"simple script miss", simpleScript, 5, 10, 3, 9, 10},
	}
	g := &Game{
		entityCollection: EntityCollection{
			Skills: skillMap{
				"Basic": charSkillMap{},
			},
		},
	}

	for _, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("tests %v", tt.name), func(t *testing.T) {
			b := &creature{Name: "Basic", Script: tt.script, Position: endless.Position_Right}
			err := b.setup()
			if err != nil {
				t.Fatalf("unable to initialize creature: %v", err)
			}
			cr, err := b.spawn(g)
			if err != nil {
				t.Fatalf("unable to spawn active creature: %v", err)
			}
			cr.CurrentVitality = tt.health
			cr.Evasion = tt.evasion

			cr.takeDamage(tt.amount, tt.accuracy)

			if cr.CurrentVitality != tt.expectHealth {
				t.Errorf("expected health to be '%v', got '%v'", tt.expectHealth, cr.CurrentVitality)
			}
		})
	}
}

func TestScriptArgs(t *testing.T) {
	script := `function testMe(creature, game)
creature.Modifiers.other_bonus = 4
if haveKey("cultists_killed", game.Memory) then
  game.Memory.cultists_killed  = game.Memory.cultists_killed +1
else
  game.Memory.cultists_killed = 1
end
end`

	l := lua.NewState()
	if err := l.DoString(script); err != nil {
		t.Fatalf("unable to parse lua script: %v", err)
	}

	fnCheck := func(L *lua.LState) int {
		have := L.ToString(1)
		mp := L.ToUserData(2)
		if m, ok := mp.Value.(map[string]interface{}); ok {
			_, ok = m[have]
			L.Push(lua.LBool(ok))
			return 1
		}

		L.Push(lua.LFalse)
		return 1
	}
	l.SetGlobal("haveKey", l.NewFunction(fnCheck))

	call := lua.P{
		Fn:      l.GetGlobal("testMe"),
		NRet:    0,
		Protect: true,
	}

	cr := &creature{
		Modifiers: map[string]int32{
			"dex_bonus": 2,
		},
	}

	game := &Game{Memory: map[string]interface{}{
		"rounds": 0,
		"thing":  "what",
	}}

	if err := l.CallByParam(call, luar.New(l, cr), luar.New(l, game)); err != nil {
		t.Fatalf("unable to call func: %v", err)
	}

	b := "other_bonus"
	var ex int32 = 4
	cv, ok := cr.Modifiers[b]
	if !ok {
		t.Fatalf("creature didn't get modifier '%v'", b)
	}

	if cv != ex {
		t.Errorf("bonus '%v' value not correct, expected '%#v' got '%#v'", b, ex, cv)
	}

	b = "cultists_killed"
	gv, ok := game.Memory[b]
	if !ok {
		t.Fatalf("game didn't get memory value '%v'", b)
	}

	var ex2 int = 1
	tmp, ok := gv.(float64)
	if !ok {
		t.Fatalf("game memory value not a float, got type: %T", gv)
	}
	i := int(tmp)

	if i != ex2 {
		t.Errorf("wrong value in game memory, expected '%v' got '%v'", ex2, i)
	}
}

func TestCreatureAct(t *testing.T) {
	script := `
function getAction()
  waveState.MonsterData.test = 1

  targets = {}
  idx = 0

  keys = entityKeys()
  for i,l in pairs(keys) do
    if l ~= creature.Id then
      targets[idx] = l
      idx = idx+1
    end
  end

  return "skill-2", targets
end`

	skScript := `
function activate()
  print("skill activated")
end`

	sk := skill{skillConfig: skillConfig{Name: "test"}, script: skScript, Level: 1}
	if err := sk.init(); err != nil {
		t.Fatalf("unable to initialize skill: %v", err)
	}

	g := &Game{
		entityCollection: EntityCollection{
			Skills: skillMap{
				"Basic": charSkillMap{
					"skill-2": &sk,
				},
			},
		},
	}

	m1 := makeTestMonster(t, g, "Basic", "c9864b71-11fa-44ec-a54b-0a3c8189e370", script, endless.Position_Right)
	m2 := makeTestMonster(t, g, "Basic", "a204ca97-bb3d-43d4-8830-1773543f5eac", "", endless.Position_Right)
	m3 := makeTestMonster(t, g, "Basic", "0fbc9de9-7aac-4811-b633-c0ff81b9bb30", "", endless.Position_Right)

	ws := newWaveState()
	ws.Entities[m1.Id] = m1
	ws.Entities[m2.Id] = m2
	ws.Entities[m3.Id] = m3

	act := m1.act(ws)

	expectTgts := []string{m2.Id, m3.Id}
	tgts := act.targets()

	if !stringSliceEq(expectTgts, tgts) {
		t.Errorf("wrong targets, expected %#v, got %#v", expectTgts, tgts)
	}
}

func TestCreatureActTarget(t *testing.T) {
	ws := newWaveState()

	g := &Game{}

	sksc := `function activate(from, to)
  to.CurrentVitality = to.CurrentVitality - 1
end`

	sk := makeTestSkill(t, g, sksc, 1)
	// m1.creature.Skills = charSkillMap{
	// 	"skill1": sk,
	// }

	g.entityCollection = EntityCollection{
		Skills: skillMap{
			"Monster 1": charSkillMap{
				"skill-2": sk,
			},
			"Monster 2": charSkillMap{
				"skill-3": sk,
			},
		},
	}

	ms1 := `function getAction()
  keys = entityKeys()

  targets = {}
  idx = 0
  for i,l in pairs(keys) do
    if l ~= creature.Id then
      targets[idx] = l
      idx = idx+1
    end
  end

  return "skill-2", targets
end`

	mid1 := "monster-1"
	m1 := makeTestMonster(t, g, "Monster 1", mid1, ms1, endless.Position_Right)

	mid2 := "monster-2"

	m2 := makeTestMonster(t, g, "Monster 2", mid2, "", endless.Position_Right)

	ws.Entities[mid1] = m1
	ws.Entities[mid2] = m2

	am := m1.act(ws)

	st := reflect.TypeOf(runSkill{})
	at := reflect.TypeOf(am)
	if st.String() != at.String() {
		t.Fatalf("wrong action message type, expected '%v', got '%v'", st.String(), at.String())
	}

	tag := am.targets()
	ex := []string{mid2}

	if !stringSliceEq(tag, ex) {
		t.Fatalf("wrong targets, expected '%#v', got '%#v'", ex, tag)
	}

	if err := am.apply(m1.creature, m2.creature, g); err != nil {
		t.Fatalf("unable to apply skill to target: %v", err)
	}

	var exH int32 = 4
	if cH := m2.creature.CurrentVitality; cH != exH {
		t.Errorf("wrong health, expected %v got %v", exH, cH)
	}
}

func makeTestPlayer(t *testing.T, g *Game, ct endless.ClassType, id string, pos endless.Position) *player {
	t.Helper()

	p, err := g.entityCollection.Classes.createPlayerClass(id, ct, g)
	if err != nil {
		t.Fatalf("unable to create player ai: %v", err)
	}
	p.is_ai = true
	return p
}

func makeTestPlayerAI(t *testing.T, g *Game, ct endless.ClassType, id, script string, pos endless.Position) *player_ai {
	t.Helper()

	p, err := g.entityCollection.Classes.createAI(id, ct, g)
	if err != nil {
		t.Fatalf("unable to create player ai: %v", err)
	}
	return p
}

func makeTestMonster(t *testing.T, g *Game, name, id, script string, pos endless.Position) *monster {
	t.Helper()

	mb := monsterBase{
		Name:   name,
		Script: script,
	}

	mt, err := createMonster(id, mb)
	if err != nil {
		t.Fatalf("unable to create monster: %v", err)
	}

	m, err := mt.spawn(g)
	if err != nil {
		t.Fatalf("unable to spawn monster: %v", err)
	}
	m.creature.Position = pos
	return m
}

func makeTestSkill(t *testing.T, g *Game, script string, level int) *skill {
	st := &skill{
		Level:  level,
		script: script,
	}
	if err := st.init(); err != nil {
		t.Fatalf("unable to initialize skill: %v", err)
	}

	sk, err := st.spawn(g)
	if err != nil {
		t.Fatalf("unable to spawn skill: %v", err)
	}
	return sk
}
