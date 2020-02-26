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
		{"no initiative", "initBase = 0", 1},
		{"simple intiative", `function initiative() return 3 end`, 3},
		{"complex initiative", `function initiative() return 11 + creature.Initiative end`, 8},
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

			init := cr.iniative()
			if init != tt.expectInit {
				t.Errorf("invalid value, expected '%v' got '%v'", tt.expectInit, init)
			}
		})
	}
}
