package game

import (
	"context"
	"fmt"
	"reflect"

	"github.com/qmuntal/stateless"
	"github.com/seanhagen/endless_stream/backend/endless"
)

type GameState string

const (
	StateCharSelect GameState = "CharacterSelect"
	StateWave                 = "Wave"
	StateDead                 = "Dead"
	StateGameOver             = "GameOver"
	StateVictory              = "Victory"
	StateStore                = "Store"
	StateNewWave              = "NewWave"
)

type GameStateTrigger string

const (
	TriggerStartGame  = "GameStart"
	TriggerAllDead    = "AllDead"
	TriggerGameOver   = "GameOver"
	TriggerStartAgain = "StartAgain"
	TriggerVictory    = "Victory"
	TriggerContinue   = "Continue"
)

func onWave(w *endless.Wave) {
	// first wave?
	//   for each player
	//     if haven't selected a character, pick one at random, send message
	//   any characters left unassigned?
	//     create AI players and assign

	// build wave
	//   figure out points available to spend
	//   is boss level?
	//     add boss for level, deduct from points
	//   add random creatures from level

	// determine iniative for all creatures, players come before monsters at same initiative
	fmt.Printf("new wave starting: %v\n", w)
}

func getGameStateMachine() *stateless.StateMachine {
	sm := stateless.NewStateMachine(StateCharSelect)
	sm.SetTriggerParameters(TriggerStartGame, reflect.TypeOf(&endless.Wave{}))

	sm.Configure(StateCharSelect).
		Permit(TriggerStartGame, StateWave)
	//, Permit(trigger interface{}, destinationState interface{}, guards ...func(context.Context, ...interface{}) bool)

	sm.Configure(StateWave).
		Permit(TriggerAllDead, StateDead).
		Permit(TriggerVictory, StateVictory).
		OnEntry(func(ctx context.Context, args ...interface{}) error {
			if len(args) <= 0 {
				return fmt.Errorf("require wave")
			}
			w, ok := args[0].(*endless.Wave)
			if !ok {
				return fmt.Errorf("first argument must be *endless.Wave")
			}
			onWave(w)

			return nil
		})

	sm.Configure(StateDead).
		Permit(TriggerStartAgain, StateCharSelect).
		Permit(TriggerGameOver, StateGameOver)

	sm.Configure(StateVictory).
		Permit(TriggerContinue, StateStore)

	sm.Configure(StateStore).
		Permit(TriggerContinue, StateNewWave)

	sm.Configure(StateNewWave).
		Permit(TriggerContinue, StateWave)

	fmt.Printf("state machine: \n\n%v\n\n", sm.ToGraph())

	fmt.Printf("state machine setup!\n")
	return sm
}
