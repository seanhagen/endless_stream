package game

import (
	"context"
	"log"

	"github.com/gofrs/uuid"
	"github.com/seanhagen/endless_stream/backend/endless"
)

// setupNewWave ...
func (g *Game) setupNewWave(ctx context.Context) error {
	// don't have 4 PCs yet, so this is the very first wave
	// need to create PCs and assign them to each player
	if len(g.playerCharacters) != 4 {
		// if player hasn't selected a class, assign random from classes left over
		for id := range g.playerIds {
			sc, ok := g.selectedCharacters[id]
			if !ok {
				sc = getNext(g.characters)
				g.selectedCharacters[id] = sc
				g.characters[sc] = id
			}
		}

		for id, class := range g.selectedCharacters {
			p, err := g.entityCollection.Classes.createPlayerClass(id, class, g)
			if err != nil {
				log.Printf("unable to create player character: %v", err)
				return err
			}
			g.playerCharacters[id] = p
		}

		// if there are fewer than 4 players, create AI players to fill the slots
		if len(g.playerCharacters) < 4 {
			for i := 4 - len(g.playerCharacters); i > 0; i-- {
				log.Printf("need to create AI player %v", i)

				id, err := uuid.NewV4()
				if err != nil {
					log.Printf("Unable to generate AI uuid: %v", err)
					return err
				}
				pid := id.String()
				sc := getNext(g.characters)
				p, err := g.entityCollection.Classes.createAI(pid, sc, g)
				if err != nil {
					log.Printf("unable to create AI player character: %v", err)
					return err
				}
				g.aiPlayers[pid] = p
			}
		}
	}

	// set up next wave
	// run all round counters

	return nil
}

// getNext returns the next availabile class, in this order:
//   1. fighter
//   2. cleric
//   3. ranger
//   4. wizard
func getNext(taken map[endless.ClassType]string) endless.ClassType {
	types := []endless.ClassType{
		endless.ClassType_Fighter,
		endless.ClassType_Cleric,
		endless.ClassType_Ranger,
		endless.ClassType_Wizard,
	}

	for _, v := range types {
		if _, ok := taken[v]; !ok {
			return v
		}
	}

	// shouldn't get here
	return endless.ClassType_Audience
}
