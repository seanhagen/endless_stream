package game

import (
	"fmt"
	"testing"

	"github.com/seanhagen/endless_stream/backend/endless"
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

			cr.init()
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

	for _, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", tt.name), func(t *testing.T) {
			p := endless.Position_Right
			b := &creature{Script: tt.script, Position: &p}
			err := b.init()
			if err != nil {
				t.Fatalf("unable to initialize creature: %v", err)
			}

			cr, err := b.spawn()
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

	for _, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("tests %v", tt.name), func(t *testing.T) {
			p := endless.Position_Right
			b := &creature{Script: tt.script, Position: &p}
			err := b.init()
			if err != nil {
				t.Fatalf("unable to initialize creature: %v", err)
			}
			cr, err := b.spawn()
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
