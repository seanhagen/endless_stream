package game

import (
	"fmt"
	"testing"
)

func TestScriptFuncExist(t *testing.T) {
	tests := []struct {
		script string
		valid  bool
	}{
		{`init = 0`, false},
		{`function init() count = 0 end`, false},
		{`function init() count = 0 end
function tick() count = 0 end`, true},
	}

	c := &creature{}
	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			s := Status{script: tt.script}
			err := s.build(c)
			if tt.valid && err != nil {
				t.Fatalf("expected valid status, got error: %v", err)
			}
			if !tt.valid && err == nil {
				t.Fatalf("expected error got nil")
			}
		})
	}
}

func TestStatusScripting(t *testing.T) {
	var startVitality int32 = 50
	var endVitality int32 = 47
	expectTicks := 3

	c := creature{
		CurrentVitality: startVitality,
	}

	// an example of a poision script that does 1 damage a turn for 3 turns
	testScript := `count = 0

function init() count = 3 end

function tick()
   count = count-1
   p = 1
   creature.CurrentVitality = creature.CurrentVitality - p
   if count == 0 then
      return true
   end
   return false
end
`
	s := Status{script: testScript}
	err := s.build(&c)
	if err != nil {
		t.Fatalf("unable to setup status: %v", err)
	}

	cnt := 0
	for {
		b, err := s.tick()
		if err != nil {
			t.Fatalf("unable to tick status: %v", err)
		}
		cnt++
		if b {
			break
		}
	}

	if c.CurrentVitality != endVitality {
		t.Errorf("wrong vitality, expected '%v' got '%v'", endVitality, c.CurrentVitality)
	}

	if cnt != expectTicks {
		t.Errorf("wrong number of ticks, expected %v got %v", expectTicks, cnt)
	}
}
