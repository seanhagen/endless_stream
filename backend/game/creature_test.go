package game

import (
	"fmt"
	"testing"
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
