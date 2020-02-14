package game

import "github.com/seanhagen/endless_stream/backend/endless"

// stateCharSelect ...
func (g *Game) stateCharSelect(input map[string][]input) error {
	startingGame := false

	// map showing who's selected what, so that we can report back to the clients
	out := map[string]endless.ClassType{}

	for pid, inputs := range input {
		if pid == g.vipPlayer {
			// check to see if there are any 'start game' inputs
			for _, i := range inputs {
				if x := i.in.GetGameStart(); x != nil {
					startingGame = true
					break
				}
			}
		}

		// check to see if any of the inputs are 'character select'
		for _, i := range inputs {
			if x := i.in.GetCharSelect(); x != nil {
				// got a character select input, x.GetChoice() -> nil or class
				//
				// if nil, player has 'unselected' the class, so remove their selection
				// otherwise:
				//   if the class is not taken, assign it to this player
				//   if the player had previously selected a class, delete that from the selected character map
			}
		}
	}

	if len(out) > 0 {
		g.output <- &endless.Output{
			Data: &endless.Output_Selected{
				Selected: &endless.CharacterSelected{
					Selected: out,
				},
			},
		}
	}

	if startingGame {
		// do the state transition!
	}

	return nil
}
