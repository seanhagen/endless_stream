package game

import (
	"log"
)

// stateCharSelect ...
func (g *Game) stateCharSelect(input map[string][]input) error {
	for pid, inputs := range input {
		for _, i := range inputs {
			// check to see if there are any 'start game' inputs from the VIP player
			if pid == g.vipPlayer && !g.started {
				if x := i.in.GetGameStart(); x != nil {
					log.Printf("VIP has started the game")
					g.started = true
					break
				}
			}

			// check to see if any of the inputs are 'character select'
			if x := i.in.GetCharSelect(); x != nil {
				c := x.GetChoice()
				if c == nil {
					delete(g.selectedCharacters, pid)

					// if nil, player has 'unselected' the class, so remove their selection
					for k, v := range g.characters {
						if v == pid {
							log.Printf("Player %v has un-selected their previous choice", pid)
							delete(g.characters, k)
							break
						}
					}
					continue
				}

				// got a character select input, x.GetChoice() -> nil or class
				ct := c.GetClass()
				if _, ok := g.characters[ct]; !ok {
					//   if the player had previously selected a class, delete that from the selected character map
					for k, v := range g.characters {
						if v == pid {
							delete(g.characters, k)
						}
					}
					//   if the class is not taken, assign it to this player
					log.Printf("Player %v selected a character: %v", pid, c)
					g.characters[ct] = pid
					g.selectedCharacters[pid] = ct
				}
			}
		}
	}

	if g.started {
		// do the state transition!
		c, ok := g.outputCountdowns["startGameCountdown"]
		if !ok {
			c = tick30Seconds + 1
		}

		c--
		if c == 0 {
			delete(g.outputCountdowns, "startGameCountdown")
			g.screenState.Fire(TriggerStartGame)
			return nil
		}
		g.outputCountdowns["startGameCountdown"] = c
	}

	return nil
}
