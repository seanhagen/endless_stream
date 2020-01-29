package game

import (
	"log"
	"sync"
	"time"

	"github.com/seanhagen/endless_stream/backend/grpc"
)

const tickLen = time.Second * 5

type player struct{}

type monster interface {
	runScript(*gameState) error
}

type gameState struct {
	initiative int

	players map[string]player

	monsters map[string]monster
}

// tick ...
func (gs *gameState) tick(t time.Time) error {
	log.Printf("game state ticking onwards")
	return nil
}

type Game struct {
	code string

	// when output needs to be sent to the clients, it's sent to this channel
	// and a listener will take care of sending the output to all connected clients
	output chan *grpc.Output

	// input from any player will be sent into this channel
	input chan *grpc.Input

	state *gameState

	players   map[chan *grpc.Output]bool
	playerIds map[string]bool

	closingClients chan chan *grpc.Output
	newClients     chan chan *grpc.Output

	lock *sync.Mutex
}

func createGame(id string) (*Game, error) {
	g := &Game{
		code: id,

		output: make(chan *grpc.Output, 1),
		input:  make(chan *grpc.Input, 1),

		state: &gameState{},

		players:   map[chan *grpc.Output]bool{},
		playerIds: map[string]bool{},

		closingClients: make(chan chan *grpc.Output),
		newClients:     make(chan chan *grpc.Output),
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
				c <- update
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
	// id, err := uuid.NewV4()
	// if err != nil {
	// 	return err
	// }

	// wg := &sync.WaitGroup{}
	// wg.Add(2)

	output := make(chan *grpc.Output)
	g.newClients <- output

	erCh := make(chan error)

	go func() {
		for {
			select {
			case ctx := <-stream.Context().Done():
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
			g.input <- msg
		}
	}()

	go func() {
		// handle outgoing messages
		for {
			out := <-output
			err := stream.Send(out)
			if err != nil {
				log.Printf("Game %v unable to send message to client: %v", g.code)
				erCh <- err
				break
			}
		}
	}()

	// wg.Wait()
	err := <-erCh
	if err != nil {
		log.Printf("error during client comms: %v", err)
	}
	return err
}
