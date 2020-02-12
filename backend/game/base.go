package game

import (
	"context"
	"sync"
	"time"

	"github.com/qmuntal/stateless"
	"github.com/seanhagen/endless_stream/backend/endless"
)

const tickLen = time.Second * 1
const stateLen = tickLen * 30

type output struct {
	id       string
	out      chan *endless.Output
	isPlayer bool
}

type input struct {
	in       *endless.Input
	isPlayer bool
}

type Game struct {
	// ctx is a cancelable context that is canceled when the game is done/erorrs
	ctx context.Context

	// code is the code to join this game
	code string

	// output is for when data needs to be sent to the clients, it's sent to this channel
	// and a listener will take care of sending the output to all connected clients
	output chan *endless.Output

	// input is a channel for data from any player
	input chan input

	// characters is a map of Class -> player ID. If a charcter hasn't been selected yet,
	// the class won't be a key in the map.
	characters map[endless.Class]string

	vipPlayer string

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

	display endless.Level

	screenState *stateless.StateMachine

	waveStateMachine  *stateless.StateMachine
	currentWaveNumber int
	waves             map[int]*endless.Wave

	// audience are humans who've connected to this game but are only able to do audience things, they
	// are unable to control any of the characters
	audience map[output]bool

	closingClients chan output
	newClients     chan output

	// lock is to protect against datat races for the maps
	lock *sync.Mutex

	// idleTime how long the game has been waiting for input, ticks up on each game tick
	idleTime int

	msgId int32

	// running is true if the game is running, or false if it's not
	// if a game isn't running, when a client connects it should
	//   - send the state
	//   - send a message stating the game is over
	//   - disconnect the client
	running bool
}

func createGame(ctx context.Context, id string) (*Game, error) {
	g := &Game{
		ctx:   ctx,
		code:  id,
		msgId: 0,

		output: make(chan *endless.Output, 10),
		input:  make(chan input, 100),

		players:          map[output]bool{},
		playerIds:        map[string]int{},
		playerCharacters: map[string]*player{},
		playerNames:      map[string]string{},

		audience: map[output]bool{},

		closingClients: make(chan output),
		newClients:     make(chan output),

		lock: &sync.Mutex{},

		screenState: getGameStateMachine(),

		display: endless.Level_Forest,

		idleTime: 0,
		running:  true,

		currentWaveNumber: 1,
		waves:             setupWaves(),
	}
	return g, nil
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
