package game

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/seanhagen/endless_stream/backend/endless"
)

const tickLen = time.Second * 1
const stateLen = tickLen * 30

type player struct{}

type monster interface {
	runScript(*gameState) error
}

type creature interface{}

type gameState struct {
	code string

	current_initiative int

	players map[string]player

	monsters map[string]monster
}

// tick ...
func (gs *gameState) tick(t time.Time) error {
	log.Printf("game state %v ticking onwards", gs.code)
	return nil
}

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
	ctx context.Context

	code string

	// when output needs to be sent to the clients, it's sent to this channel
	// and a listener will take care of sending the output to all connected clients
	output chan *endless.Output

	// input from any player will be sent into this channel
	input chan input

	state *gameState

	players   map[output]bool
	playerIds map[string]bool

	audience    map[output]bool
	audienceIds map[string]bool

	closingClients chan output
	newClients     chan output

	lock *sync.Mutex

	idleTime int
}

func createGame(ctx context.Context, id string) (*Game, error) {
	g := &Game{
		ctx: ctx,

		code: id,

		output: make(chan *endless.Output, 10),
		input:  make(chan input, 100),

		state: &gameState{
			code:     id,
			players:  map[string]player{},
			monsters: map[string]monster{},
		},

		players:     map[output]bool{},
		playerIds:   map[string]bool{},
		audienceIds: map[string]bool{},

		closingClients: make(chan output),
		newClients:     make(chan output),
	}
	return g, nil
}
