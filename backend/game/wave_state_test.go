package game

import (
	"testing"
)

func TestWaveStateTick(t *testing.T) {
	ws := newWaveState()
	mid := "1"
	m, err := createMonster(mid, monsterBase{Name: "Rat"})
	if err != nil {
		t.Fatalf("unable to create monster: %v", err)
	}

	ci := ws.current_initiative
	ws.initiative[ci] = []actor{m}
	ws.Entities[mid] = m
	// m1 := makeCreature(t, g, mid, script, p)
	// m2 := makeCreature(t, g, mid2, script, p)
	// ws := newWaveState()
	// ws.Entities[mid] = m1
	// ws.Entities[mid2] = m2

	t.Errorf("not yet")
}
