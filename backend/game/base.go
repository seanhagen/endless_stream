package game

import (
	"context"
	"sync"
	"time"

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

	// state the actual state of the game, what players are human controlled, what the current
	// wave is, what monsters are in this wave, player xp & hp levels, etc
	state *gameState

	// players is a map of output -> bool, to keep track of connected players
	players map[output]bool
	// playerIds is a map of string -> int, meant to keep track of player ids
	//   when a player connects for the first time:
	//     the game generates an id, stores it as a key in this map with a value of 1
	//   when a player disconnects, the value is decremented ( so it should be 0, but may be lower if things get real fucky )
	//   when a player re-connects, if they send a register with the id they had previously
	//     if that id is in the map and the value is < 1, they can reclaim their player
	//     if the id isn't in the map or the value is >= 1, they're assigned to the audience
	playerIds map[string]int

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

		state: newState(ctx, id),

		players:  map[output]bool{},
		audience: map[output]bool{},

		playerIds: map[string]int{},

		closingClients: make(chan output),
		newClients:     make(chan output),

		lock: &sync.Mutex{},

		idleTime: 0,
		running:  true,
	}
	return g, nil
}
