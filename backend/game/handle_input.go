package game

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/seanhagen/endless_stream/backend/endless"
)

// handleInput ...
func (g *Game) handleInput(ctx context.Context, in input) error {
	g.lock.Lock()
	defer g.lock.Unlock()

	var err error
	var out *endless.Output

	if in.isPlayer {
		out, err = g.handlePlayerInput(ctx, in.in)
	} else {
		out, err = g.handleAudienceInput(ctx, in.in)
	}

	if err != nil {
		g.lock.Lock()
		defer g.lock.Unlock()
		mid := g.msgId
		g.msgId++

		g.output <- &endless.Output{
			Data: &endless.Output_Msg{
				Msg: &endless.EventMessage{
					MsgId: mid,
					// TODO: error -> json string? need error messages that make sense to the players?
					Msg:      fmt.Sprintf("Unable to handle input: %v", err),
					IsError:  true,
					PlayerId: &wrappers.StringValue{Value: in.in.GetPlayerId()},
				},
			},
		}
		return err
	}

	if out != nil {
		g.output <- out
	}

	return nil
}

// handlePlayerInput ...
func (g *Game) handlePlayerInput(ctx context.Context, in *endless.Input) (*endless.Output, error) {
	id := in.GetPlayerId()

	switch g.screenState.MustState().(GameState) {
	case StateCharSelect:
		// character select
		if cs := in.GetCharSelect(); cs != nil {
			return g.handleCharacterSelect(id, cs)
		}
		if st := in.GetGameStart(); st != nil {
			return g.handleGameStart(id, st)
		}

	case StateWave:
		id := in.GetPlayerId()
		if g.isCurrentPlayer(id) {
			// handled by wave
			// in.GetActionComplete()
			// in.GetItem()
			// in.GetMove()
			// in.GetSkill()
		}
	case StateDefeat:
		// game over
		// cont := in.GetContinue() // start new game, go to the
		// eg := in.GetEndGame()
	case StateVictory:
		// cont := in.GetContinue()
	case StateStore:
		// handled by store
		// pur := in.GetPurchase()
	case StateNewWave:
		// cont := in.GetContinue()
	case StateGameOver:
		// no input taken here, game is over
	}

	return nil, nil
}

// handleAudienceInput ...
func (g *Game) handleAudienceInput(ctx context.Context, in *endless.Input) (*endless.Output, error) {
	return nil, nil
}

// handleCharacterSelect ...
func (g *Game) handleCharacterSelect(id string, in *endless.CharSelect) (*endless.Output, error) {
	s := g.screenState.MustState().(GameState)
	if s != StateCharSelect {
		id := in.GetPlayerId()
		mid := g.msgId
		g.msgId++
		return &endless.Output{
			Data: &endless.Output_Msg{
				Msg: &endless.EventMessage{
					MsgId:    mid,
					Msg:      "Can't select character right now",
					IsError:  true,
					IsAlert:  false,
					PlayerId: &wrappers.StringValue{Value: id},
					LogOnly:  false,
				},
			},
		}, nil
	}

	c := in.GetChoice()
	// _, err := uuid.FromString(in.CharSelect.GetPlayerId())

	if _, ok := g.characters[c]; !ok {
		// nobody has taken the character yet
		g.characters[c] = id
		g.playerCharacters[id] = createPlayerCharacter(id, c)
		out := map[string]endless.Class{}
		for k, v := range g.characters {
			out[v] = k
		}

		return &endless.Output{
			Data: &endless.Output_Selected{
				Selected: &endless.CharacterSelected{
					Selected: out,
				},
			},
		}, nil
	}

	return nil, nil
}

// handleGameStart ...
func (g *Game) handleGameStart(id string, in *endless.GameStart) (*endless.Output, error) {
	if id != g.vipPlayer {
		mid := g.msgId
		g.msgId++
		return &endless.Output{
			Data: &endless.Output_Msg{
				Msg: &endless.EventMessage{
					MsgId:    mid,
					Msg:      "Only the VIP can start the game",
					IsError:  true,
					PlayerId: &wrappers.StringValue{Value: id},
				},
			},
		}, nil
	}

	err := g.screenState.Fire(TriggerStartGame)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

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
