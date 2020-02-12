package game

import (
	"fmt"
	"reflect"

	"github.com/qmuntal/stateless"
	"github.com/seanhagen/endless_stream/backend/endless"
)

// GameState ...
type GameState string

const (
	StateCharSelect   GameState = "CharacterSelect"
	StateNewWave                = "NewWave"
	StateWave                   = "Wave"
	StateWaveInput              = "WaveInput"
	StateWaveProcess            = "WaveProcess"
	StateWaveAnimWait           = "WaveAnimWait"
	StateDefeat                 = "Defeat"
	StateVictory                = "Victory"
	StateStore                  = "Store"
	StateGameOver               = "GameOver"
)

type GameStateTrigger string

const (
	// Char Select -> New Wave
	TriggerStartGame GameStateTrigger = "GameStart"

	// New Wave -> Wave
	TriggerStartWave = "StartWave"

	// Wave -> Wave Input
	TriggerWaveInput = "GetWaveInput"

	// Wave -> Victory
	TriggerWaveVictory = "WaveVictory"

	// Wave -> Defeat
	TriggerWaveDefeat = "WaveDefeat"

	// Wave Input -> Wave Process
	TriggerWaveProcessing = "WaveProcessing"

	// Wave Process -> Anim Wait
	TriggerWaveAnimWait = "WaveAnimWait"

	// Anim Wait -> Wave
	TriggerWaveContinue = "WaveContinue"

	// Defeat -> Char Select
	TriggerNewGame = "StartOver"

	// Defeat -> Game Over
	TriggerGameOver = "IGiveUp"

	// Victory -> Store
	TriggerStore = "LetsGoShopping"

	// Store -> New Wave
	TriggerNewWave = "BackIntoTheBreach"
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
		Permit(TriggerStartGame, StateNewWave)
	//, Permit(
	//    trigger interface{},
	//    destinationState interface{},
	//    guards ...func(context.Context, ...interface{}) bool)

	sm.Configure(StateNewWave).
		Permit(TriggerStartWave, StateWave)

	sm.Configure(StateWave).
		Permit(TriggerWaveInput, StateWaveInput).
		Permit(TriggerWaveVictory, StateVictory).
		Permit(TriggerWaveDefeat, StateDefeat)

	sm.Configure(StateWaveInput).
		Permit(TriggerWaveProcessing, StateWaveProcess)

	sm.Configure(StateWaveProcess).
		Permit(TriggerWaveAnimWait, StateWaveAnimWait)

	sm.Configure(StateWaveAnimWait).
		Permit(TriggerWaveContinue, StateWave)

	sm.Configure(StateDefeat).
		Permit(TriggerNewGame, StateCharSelect).
		Permit(TriggerGameOver, StateGameOver)

	sm.Configure(StateVictory).
		Permit(TriggerStore, StateStore)

	sm.Configure(StateStore).
		Permit(TriggerNewWave, StateNewWave)

	// sm.Configure(StateWave).
	// 	OnEntry(func(ctx context.Context, args ...interface{}) error {
	// 		if len(args) <= 0 {
	// 			return fmt.Errorf("require wave")
	// 		}
	// 		w, ok := args[0].(*endless.Wave)
	// 		if !ok {
	// 			return fmt.Errorf("first argument must be *endless.Wave")
	// 		}
	// 		onWave(w)
	// 		return nil
	// 	})

	// fmt.Printf("state machine: \n\n%v\n\n", sm.ToGraph())
	// fmt.Printf("state machine setup!\n")
	return sm
}
