package game

// func TestGameStateSetup(t *testing.T) {
// 	s := getGameStateMachine()

// 	fmt.Printf("graph: \n\n%v\n\n", s.ToGraph())

// 	// err := s.Fire(TriggerAllDead, "wat")
// 	// spew.Dump(s, err)

// 	w := &endless.Wave{
// 		Num:     1,
// 		HasBoss: false,
// 		Level:   endless.Level_Forest,
// 	}

// 	err := s.Fire(TriggerStartGame, w)
// 	spew.Dump(err)

// 	fmt.Printf("state is: %v\n", s.MustState())

// 	t.Errorf("nope")
// }
