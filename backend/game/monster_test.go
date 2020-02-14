package game

import (
	"fmt"
	"testing"
)

func TestMonsterInit(t *testing.T) {
	tests := []struct {
		name string

		str int32
		int int32
		agi int32

		xp   int32
		gold int32
		cost int32

		mod     int32
		cmod    int32
		goldMod float64
		xpMod   float64

		isBoss bool
	}{
		{"rat", 0, 0, 1, 1, 1, 2, 6, -1, 2.0, 1.0, false},
		{"corvid shaman", 2, 4, 3, 2, 4, 8, 6, 2, 2, 1, false},
		{"wendigo", 6, 5, 7, 16, 32, 16, 12, 5, 2, 1, true},
		{"leprechaun", 1, 0, 7, 7, 28, 7, 6, 0, 4.0, 1.0, true},
	}

	for _, x := range tests {
		tt := x
		tn := fmt.Sprintf("test %v", tt.name)
		t.Run(tn, func(t *testing.T) {
			mb := monsterBase{
				Name:         tt.name,
				CostMod:      tt.cmod,
				Mod:          tt.mod,
				GoldMod:      tt.goldMod,
				XPMod:        tt.xpMod,
				Strength:     tt.str,
				Intelligence: tt.int,
				Agility:      tt.agi,
				IsBoss:       tt.isBoss,
			}

			monster := createMonster(tn, mb, "")

			if monster.Gold != tt.gold {
				t.Errorf("Wrong gold amount, expected %v got %v", tt.gold, monster.Gold)
			}

			if monster.XP != tt.xp {
				t.Errorf("Wrong XP amount, expected %v got %v", tt.xp, monster.XP)
			}

			if monster.cost != tt.cost {
				t.Errorf("Wrong cost amount, expected %v got %v", tt.cost, monster.cost)
			}
		})
	}
}
