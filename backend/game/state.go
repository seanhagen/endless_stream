package game

import (
	"context"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	"github.com/qmuntal/stateless"
	"github.com/seanhagen/endless_stream/backend/endless"
)

type skill interface{}

type item interface{}

type status interface {
	effect(*creature)
	tick()
	shouldRemove() bool
}

type creature struct {
	id   uuid.UUID
	name string

	position *endless.Position
	statuses []status

	strength         int32
	maxVitality      int32 // aka max hp
	currentVitality  int32 // aka current hp
	combatDamageBase int32
	vitalityRegen    int32

	intelligence int32
	currentFocus int32
	maxFocus     int32
	willpower    int32
	focusRegen   int32

	agility    int32
	evasion    int32
	accuracy   int32
	initiative int32

	// a list of current modifiers that affect the various stats
	modifiers map[string]int32

	gold int32
	xp   int32
}

type player struct {
	creature

	class endless.Class

	isAI  bool
	level int32

	skills    map[string]skill
	inventory map[string]item
}

type monster struct {
	creature
	mType    endless.Type
	isFlying bool
	isBoss   bool
	cost     int32
}

type actor interface {
	isPlayer() bool
	id() uuid.UUID
}

type gameState struct {
	code string

	monsters []monster

	// players is a map of UUID -> player structs
	players map[string]player
	// playerNames is a map of UUID -> player name
	playerNames map[string]string

	currentPlayer *uuid.UUID

	// lock is to prevent race conditions
	lock *sync.Mutex

	// screenState is a state machine that controls what screen the game is on
	screenState *stateless.StateMachine

	// what "level" ( ie, forest, cave, dungeon, etc ) should the UI display
	level endless.Level

	// waveNumber ...
	waveNumber int
	// currentWave ...
	currentWave *waveState
	// waveState is a state machine that handles what happens within a wave
	waveState *stateless.StateMachine
}

type waveState struct {
	current_initiative_step int
	current_initiative      int
	max_initiative          int

	initiative map[int][]actor

	info endless.Wave

	// monsterData is a map of string -> data that is passed in when a creature
	// performs various actions
	//
	// For example, when a cultist dies it increments a counter -- when the counter hits 7, a shoggoth is summoned.
	//
	// That information is stored here.
	monsterData map[string]interface{}
}

func newState(ctx context.Context, id string) *gameState {
	return &gameState{
		code: id,
		// current_initiative_step: 0,
		// current_initiative:      0,
		// max_initiative:          0,
		// players:                 map[string]player{},
		// audience:                map[string]audience{},
		// monsters:                map[string]monster{},
	}
}

// tick ...
func (gs *gameState) tick(t time.Time) error {
	// get current actor

	// actor has action to perform?
	//   check for statuses like stunned
	//   if need to skip, return a special "skip" action

	// if no action ( ie, still waiting for player input )

	// perform action
	//

	return nil
}

// handleInput ...
func (gs *gameState) handleInput(ctx context.Context, in input) error {
	if in.isPlayer {
		return gs.handlePlayerInput(ctx, in.in)
	}
	return gs.handleAudienceInput(ctx, in.in)
}

// handleAudienceInput ...
func (gs *gameState) handleAudienceInput(ctx context.Context, in *endless.Input) error {
	return nil
}

// handlePlayerInput ...
func (gs *gameState) handlePlayerInput(ctx context.Context, in *endless.Input) error {
	return nil
}
