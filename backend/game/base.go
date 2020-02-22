package game

import (
	"context"
	"fmt"
	"sync"
	"time"

	packr "github.com/gobuffalo/packr/v2"
	"github.com/qmuntal/stateless"
	"github.com/seanhagen/endless_stream/backend/endless"
)

// tick 5 times a second
const tickLen = time.Millisecond * 200

// send state every 5 seconds
const stateLen = time.Second * 5

// max number of rounds ahead something can schedule: 30
const roundCounterMax int32 = 30
const roundCounterMin int32 = 0

// max number of ticks ahead something can schedule: 60 seconds, or 300 ticks
const tickCounterMax int32 = 300
const tickCounterMin int32 = 0

// 30 seconds in ticks is 150
const tick30Seconds int32 = 150

const inputLength = 100
const audienceInputLength = inputLength

type output struct {
	id         string
	out        chan *endless.Output
	isPlayer   bool
	isAudience bool
	isDisplay  bool
}

type input struct {
	in       *endless.Input
	isPlayer bool
}

type countdownFunc func(context.Context)

type Box interface {
	Find(string) ([]byte, error)
	FindString(string) (string, error)
	HasDir(string) bool
	Has(string) bool
	Walk(packr.WalkFunc) error
	WalkPrefix(string, packr.WalkFunc) error
}

type Game struct {
	// ctx is a cancelable context that is canceled when the game is done/erorrs
	ctx context.Context

	cancelFn context.CancelFunc

	// code is the code to join this game
	code string

	// output is for when data needs to be sent to the clients, it's sent to this channel
	// and a listener will take care of sending the output to all connected clients
	output chan *endless.Output

	// input is a channel for data from any player or audience member
	input chan input

	// playerInput is input just from the 1-4 players, no audience input
	playerInput chan input

	// audienceInput is input just from the audience, no player input
	audienceInput chan input

	// characters is a map of Class -> player ID. If a charcter hasn't been selected yet,
	// the class won't be a key in the map.
	characters map[endless.ClassType]string

	displayClients map[output]bool

	// players is a map of output -> bool, to keep track of connected players
	players map[output]bool
	// playerIds is a map of string -> int, meant to keep track of player ids
	//   when a player connects for the first time:
	//     the game generates an id, stores it as a key in this map with a value of 1
	//   when a player disconnects, the value is decremented
	//     ( so it should be 0, but may be lower if things get real fucky )
	//   when a player re-connects, if they send a register with the id they had previously
	//     if that id is in the map and the value is < 1, they can reclaim their player
	//     if the id isn't in the map or the value is >= 1, they're assigned to the audience
	playerIds        map[string]int
	playerCharacters map[string]*player
	playerNames      map[string]string
	currentPlayer    *string

	selectedCharacters map[endless.ClassType]string

	// vipPlayer is the player who connected first
	vipPlayer string

	display endless.Level

	screenState *stateless.StateMachine

	waveStateMachine  *stateless.StateMachine
	currentWaveNumber int
	waves             map[int]*endless.Wave

	// audience are humans who've connected to this game but are only able to do audience things, they
	// are unable to control any of the characters
	audience    map[output]bool
	audienceIds map[string]bool
	// they do get to cheer or boo, which can impact the game ( ...somehow )
	audienceCheer int
	audienceBoo   int

	// lock is to protect against datat races for the maps
	lock *sync.Mutex

	// idleTime how long the game has been waiting for input, ticks up on each game tick
	idleTime  int
	idleSince time.Time
	connected int

	msgId int32

	// round countdowns are functions that wait a specific number of rounds before executing
	roundCounterIdx int32
	roundCountdowns map[int32][]countdownFunc

	// tick countdowns are functions that wait a specific number of ticks before executing
	tickCounterIdx int32
	tickCountdowns map[int32][]countdownFunc

	// Running is true if the game is running, or false if it's not
	// if a game isn't running, when a client connects it should
	//   - send the state
	//   - send a message stating the game is over
	//   - disconnect the client
	running bool

	// Started is true if the VIP has started the game. This is to prevent people
	// from joining the game and taking over AI players.
	started bool

	entityCollection EntityCollection
}

func Create(ctx context.Context, id string, ec EntityCollection) (*Game, error) {
	ctx, cancel := context.WithCancel(ctx)
	g := &Game{
		ctx:      ctx,
		cancelFn: cancel,
		code:     id,
		msgId:    0,

		entityCollection: ec,

		output:        make(chan *endless.Output, 10),
		input:         make(chan input, inputLength),
		playerInput:   make(chan input, inputLength),
		audienceInput: make(chan input, audienceInputLength),

		displayClients: map[output]bool{},

		players:          map[output]bool{},
		playerIds:        map[string]int{},
		playerCharacters: map[string]*player{},
		playerNames:      map[string]string{},

		selectedCharacters: map[endless.ClassType]string{},

		audience: map[output]bool{},

		lock: &sync.Mutex{},

		screenState: getGameStateMachine(),

		display: endless.Level_Forest,

		idleTime: 0,
		running:  true,

		currentWaveNumber: 1,
		waves:             setupWaves(),

		roundCounterIdx: roundCounterMax,
		roundCountdowns: setupRoundCountdowns(),

		tickCounterIdx: tickCounterMax,
		tickCountdowns: setupTickCountdowns(),
	}

	return g, nil
}

// addRoundCountdown ...
func (g *Game) addRoundCountdown(numRounds int32, fn countdownFunc) error {
	return addCountdown(
		numRounds,
		roundCounterMin,
		roundCounterMax,
		g.roundCounterIdx,
		fn,
		g.roundCountdowns,
	)
}

// addTickCountdown ...
func (g *Game) addTickCountdown(numRounds int32, fn countdownFunc) error {
	return addCountdown(
		numRounds,
		tickCounterMin,
		tickCounterMax,
		g.tickCounterIdx,
		fn,
		g.tickCountdowns,
	)
}

// addCountdown ...
func addCountdown(numRounds, min, max, cur int32, fn countdownFunc, cdowns map[int32][]countdownFunc) error {
	if numRounds > max {
		return fmt.Errorf("'%v' is greater than max counter '%v'", numRounds, max)
	}

	if numRounds < min {
		return fmt.Errorf("'%v' is smaller than min counter '%v'", numRounds, min)
	}

	idx := cur - numRounds
	if idx < 0 {
		idx += max
	}

	c, ok := cdowns[idx]
	if !ok {
		c = []countdownFunc{}
	}
	c = append(c, fn)
	cdowns[idx] = c
	return nil
}

func setupRoundCountdowns() map[int32][]countdownFunc {
	return setupCountdowns(roundCounterMin, roundCounterMax)
}

func setupTickCountdowns() map[int32][]countdownFunc {
	return setupCountdowns(tickCounterMin, tickCounterMax)
}

func setupCountdowns(min, max int32) map[int32][]countdownFunc {
	out := map[int32][]countdownFunc{}
	for i := min; i <= max; i++ {
		out[i] = []countdownFunc{}
	}
	return out
}

// setupWaves ...
func setupWaves() map[int]*endless.Wave {
	out := map[int]*endless.Wave{}
	for i := 1; i <= 10; i++ {
		out[i] = createWave(i)
	}
	return out
}

// createWave ...
func createWave(i int) *endless.Wave {
	return &endless.Wave{
		Num:     int32(i),
		HasBoss: i%10 == 0,
		Level:   getLevelOfWave(i),
	}
}

func getLevelOfWave(i int) endless.Level {
	if i <= 10 {
		return endless.Level_Forest
	}

	if i > 10 && i <= 20 {
		return endless.Level_Cave
	}

	if i > 20 && i <= 30 {
		return endless.Level_Dungeon
	}

	if i > 30 && i <= 40 {
		return endless.Level_Ice
	}

	if i > 40 && i <= 50 {
		return endless.Level_Fire
	}

	return endless.Level_Void
}

// isCurrentPlayer ...
func (g *Game) isCurrentPlayer(id string) bool {
	if g.currentPlayer == nil {
		return false
	}
	return id == *g.currentPlayer
}
