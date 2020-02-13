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

		mmod int32
		gmod int32
		cmod int32

		boss bool
	}{
		{"rat", 0, 0, 1, 1, 1, 2, 6, -3, -1, false},
		{"corvid shaman", 2, 4, 3, 2, 4, 8, 6, -1, 2, false},
		{"wendigo", 6, 5, 7, 16, 32, 16, 12, 3, 5, true},
	}

	for _, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", tt.name), func(t *testing.T) {
			mb := monsterBase{
				Name:         tt.name,
				CostMod:      tt.cmod,
				MMod:         tt.mmod,
				GMod:         tt.gmod,
				Strength:     tt.str,
				Intelligence: tt.int,
				Agility:      tt.agi,
				IsBoss:       tt.boss,
			}

			monster := createMonster(mb, "")

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
