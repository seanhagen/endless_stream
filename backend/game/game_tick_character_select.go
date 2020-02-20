package game

import (
	"log"

	"github.com/davecgh/go-spew/spew"
)

// stateCharSelect ...
func (g *Game) stateCharSelect(input map[string][]input) error {
	startingGame := false
	// updateSelected := false

	if len(input) > 0 {
		spew.Dump(input)
	}

	for pid, inputs := range input {
		for _, i := range inputs {
			// check to see if there are any 'start game' inputs from the VIP player
			if pid == g.vipPlayer {
				if x := i.in.GetGameStart(); x != nil {
					startingGame = true
					break
				}
			}

			// check to see if any of the inputs are 'character select'
			if x := i.in.GetCharSelect(); x != nil {
				c := x.GetChoice()
				if c == nil {
					// if nil, player has 'unselected' the class, so remove their selection
					for k, v := range g.selectedCharacters {
						if v == pid {
							log.Printf("Player %v has un-selected their previous choice", pid)
							delete(g.selectedCharacters, k)
							// updateSelected = true
							break
						}
					}
					continue
				}

				// got a character select input, x.GetChoice() -> nil or class
				ct := c.GetClass()
				if _, ok := g.selectedCharacters[ct]; !ok {
					//   if the player had previously selected a class, delete that from the selected character map
					for k, v := range g.selectedCharacters {
						if v == pid {
							delete(g.selectedCharacters, k)
						}
					}
					//   if the class is not taken, assign it to this player
					log.Printf("Player %v selected a character: %v", pid, c)
					g.selectedCharacters[ct] = pid
					// updateSelected = true
				}
			}
		}
	}

	if startingGame {
		// do the state transition!
		log.Printf("VIP has started the game")
	}

	return nil
}
