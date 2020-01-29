package game

import (
	"log"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	"github.com/seanhagen/endless_stream/backend/grpc"
)

const tickLen = time.Second * 5

type player struct{}

type monster interface {
	runScript(*gameState) error
}

type creature interface{}

type gameState struct {
	current_initiative int

	players map[string]player

	monsters map[string]monster
}

// tick ...
func (gs *gameState) tick(t time.Time) error {
	log.Printf("game state ticking onwards")
	return nil
}

type output struct {
	id       string
	out      chan *grpc.Output
	isPlayer bool
}

type input struct {
	in       *grpc.Input
	isPlayer bool
}

type Game struct {
	code string

	// when output needs to be sent to the clients, it's sent to this channel
	// and a listener will take care of sending the output to all connected clients
	output chan *grpc.Output

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
}

func createGame(id string) (*Game, error) {
	g := &Game{
		code: id,

		output: make(chan *grpc.Output, 1),
		input:  make(chan input, 1),

		state: &gameState{},

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

	for {
		select {
		case newClient := <-g.newClients:
			g.players[newClient] = true

		case clientLeft := <-g.closingClients:
			delete(g.players, clientLeft)

		case update := <-g.output:
			log.Printf("sending output to players")
			for c := range g.players {
				log.Printf("player '%v' (isPlayer: %v)", c.id, c.isPlayer)
				c.out <- update
			}

		case input := <-g.input:
			log.Printf("got player input: %v", input)

		case tick := <-ticker.C:
			g.state.tick(tick)
		}
	}
}

// registerClient ...
func (g *Game) registerClient(stream grpc.GameServer_StateServer) error {
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

	out := &grpc.Output{
		Data: &grpc.Output_Joined{
			Joined: &grpc.JoinedGame{
				Id:       id.String(),
				IsPlayer: isPlayer,
			},
		},
	}
	stream.Send(out)

	output := output{
		id:       id.String(),
		out:      make(chan *grpc.Output),
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
				log.Printf("Game %v stream context done", g.code)
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
			msg.Id = id.String()
			g.input <- input{in: msg, isPlayer: isPlayer}
		}
	}()

	go func() {
		// handle outgoing messages
		for {
			out := <-output.out
			err := stream.Send(out)
			if err != nil {
				log.Printf("Game %v unable to send message to client: %v", g.code)
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
