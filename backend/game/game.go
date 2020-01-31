package game

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes"
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

// listen ...
func (g *Game) listen() {
	ticker := time.NewTicker(tickLen)
	stateTick := time.NewTicker(stateLen)

	for {
		select {
		case newClient := <-g.newClients:
			log.Printf("client connected: %v", newClient.id)
			g.players[newClient] = true

		case clientLeft := <-g.closingClients:
			log.Printf("client disconnected: %v", clientLeft.id)
			delete(g.players, clientLeft)

		case update := <-g.output:
			if len(g.players) > 0 {
				log.Printf("sending output to players")
				for c := range g.players {
					log.Printf("player '%v' (isPlayer: %v)", c.id, c.isPlayer)
					c.out <- update
				}
			}

		case input := <-g.input:
			log.Printf("got player input: %v", input)

		case tick := <-ticker.C:
			log.Printf("game tick")
			g.state.tick(tick)
			ts, _ := ptypes.TimestampProto(tick)

			g.output <- &endless.Output{
				Data: &endless.Output_Tick{
					Tick: &endless.Tick{Time: ts},
				},
			}

		case <-stateTick.C:
			log.Printf("sending state")
			g.output <- &endless.Output{
				Data: &endless.Output_State{
					State: &endless.CurrentState{},
				},
			}

		case <-g.ctx.Done():
			log.Printf("game context signaled done!")
			goto finished
		}
	}
finished:
	log.Printf("game done")
}

// registerClient ...
func (g *Game) registerClient(stream endless.Game_StateServer) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	isPlayer := true
	if len(g.players) == 4 {
		isPlayer = false
	}

	if isPlayer {
		g.playerIds[id.String()] = true
	} else {
		g.audienceIds[id.String()] = true
	}

	out := &endless.Output{
		Data: &endless.Output_Joined{
			Joined: &endless.JoinedGame{
				Id:         id.String(),
				AsAudience: !isPlayer,
			},
		},
	}
	stream.Send(out)

	output := output{
		id:       id.String(),
		out:      make(chan *endless.Output),
		isPlayer: isPlayer,
	}
	g.newClients <- output
	defer func() {
		g.closingClients <- output
	}()

	erCh := make(chan error)

	go func() {
		for {
			select {
			case _ = <-stream.Context().Done():
				log.Printf("Game %v stream client context done", g.code)
				erCh <- nil
			}
		}
	}()

	go func() {
		// handle incoming messages
		for {

			msg, err := stream.Recv()
			if err != nil {
				log.Printf("Game %v unable to receive message from client: %v", g.code, err)
				erCh <- err
				break
			}
			msg.PlayerId = id.String()
			g.input <- input{in: msg, isPlayer: isPlayer}
		}
	}()

	go func() {
		// handle outgoing messages
		for {
			out := <-output.out
			err := stream.Send(out)
			if err != nil {
				log.Printf("Game %v unable to send message to client: %v", g.code, err)
				erCh <- err
				break
			}
		}
	}()

	// wg.Wait()
	err = <-erCh
	if err != nil {
		log.Printf("error during client comms: %v", err)
	}
	return err
}
