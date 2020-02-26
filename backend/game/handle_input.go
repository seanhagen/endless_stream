package game

import (
	"github.com/seanhagen/endless_stream/backend/endless"
)

// handleInput ...
func (g *Game) handleInput(in input) {
	// log.Printf("got player/audience input: %#v", in)
	if in.isPlayer {
		g.playerInput <- in
	} else {
		g.audienceInput <- in
	}
}

// getPlayerInput ...
func (g *Game) getPlayerInput() map[string][]input {
	out := map[string][]input{}
	l := len(g.playerInput)
	// log.Printf("player inputs to be processed: %v", l)
	if l > 0 {
		for i := 0; i < l; i++ {
			pi := <-g.playerInput
			id := pi.in.GetPlayerId()
			ins, ok := out[id]
			if !ok {
				ins = []input{}
			}
			ins = append(ins, pi)
			out[id] = ins
		}
	}
	return out
}

// getAudienceInputs ...
func (g *Game) getAudienceInputs() []input {
	out := []input{}
	l := len(g.audienceInput)
	if l > 0 {
		for i := 0; i < l; i++ {
			ai := <-g.audienceInput
			if y := ai.in.GetAudience(); y != nil {
				out = append(out, ai)
			}
		}
	}

	return out
}

// // handlePlayerInput ...
// func (g *Game) handlePlayerInput(ctx context.Context, in *endless.Input) (*endless.Output, error) {
// 	id := in.GetPlayerId()

// 	// it's input from a player, figure out what to do based on current game state
// 	switch g.screenState.MustState().(GameState) {
// 	case StateCharSelect:
// 		// character select
// 		if cs := in.GetCharSelect(); cs != nil {
// 			return g.handleCharacterSelect(id, cs)
// 		}
// 		if st := in.GetGameStart(); st != nil {
// 			return g.handleGameStart(id, st)
// 		}

// 	case StateWaveInput:
// 		if g.isCurrentPlayer(id) {
// 			// in.GetItem()
// 			// in.GetMove()
// 			// in.GetSkill()
// 		}

// 	case StateWaveAnimWait:
// 		// in.GetActionComplete()

// 	case StateDefeat:
// 		// game over
// 		if id == g.vipPlayer {
// 			// cont := in.GetContinue() // start new game, go to the
// 			// eg := in.GetEndGame()
// 		}

// 	case StateVictory:
// 		// battle over, heroes won
// 		if id == g.vipPlayer {
// 			// cont := in.GetContinue()
// 		}

// 	case StateStore:
// 		// handled by store
// 		// pur := in.GetPurchase()
// 		if id == g.vipPlayer {
// 			// cont := in.GetContinue()
// 		}
// 	case StateNewWave:
// 		if id == g.vipPlayer {
// 			// cont := in.GetContinue()
// 		}
// 	case StateGameOver:
// 		// no input taken here, game is over
// 	}
// 	return nil, nil
// }

// createPlayerCharacter ...
func createPlayerCharacter(id string, c endless.Class) *player {
	var p endless.Position
	p = endless.Position_Left

	return &player{
		creature: creature{
			Id:        id,
			Position:  &p,
			Modifiers: map[string]int32{},
		},
		class:     c,
		isAI:      false,
		level:     1,
		skills:    map[string]skill{},
		inventory: map[string]item{},
	}
}
