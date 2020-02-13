package game

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestStatusScripting(t *testing.T) {
	c := creature{
		CurrentVitality: 50,
		MaxVitality:     50,
	}

	testScript := `count = 0

function init()
   count = math.random(1,3)
   print("status init! count: " .. count)
end

function tick()
   print("creature current health " .. creature.CurrentVitality)

   count = count-1
   p = math.random(1,5)
   print("poison damage: " .. p)
   creature.CurrentVitality = creature.CurrentVitality - p
   print("creature health after poison damage " .. creature.CurrentVitality)
   if count == 0 then
      return true
   end
   return false
end
`

	s, err := newStatus(testScript, &c)
	if err != nil {
		t.Errorf("unable to setup status: %v", err)
	}

	cnt := 0
	for {
		b, err := s.tick()
		if err != nil {
			t.Fatalf("unable to tick status: %v", err)
		}
		if b {
			break
		}
		cnt++
	}
	t.Logf("got %v ticks from status", cnt)

	spew.Dump(c.CurrentVitality)

	t.Errorf("nope")
}
