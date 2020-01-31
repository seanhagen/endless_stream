package game

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/seanhagen/endless_stream/backend/endless"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

type Srv struct {
	games   map[string]*Game
	cancels map[string]context.CancelFunc
}

func getGameId() string {
	b := make([]byte, 4)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// Create ...
func (s *Srv) Create(ctx context.Context, in *endless.CreateGame) (*endless.GameCreated, error) {
	var id string
	for {
		id = getGameId()
		if _, ok := s.games[id]; ok {
			continue
		}
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		g, err := createGame(ctx, id)
		if err != nil {
			cancel()
			return nil, err
		}
		go g.listen()
		s.cancels[id] = cancel
		s.games[id] = g
		break
	}

	return &endless.GameCreated{Code: id}, nil
}

// State ...
func (s *Srv) State(stream endless.Game_StateServer) error {
	var game *Game
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Printf("Unable to receive state message during handshake: %v", err)
			return fmt.Errorf("unable to receive message? %v", err)
		}

		r := msg.GetRegister()
		if r == nil {
			out := &endless.Output{
				Data: &endless.Output_Msg{
					Msg: &endless.EventMessage{
						Msg:     "You must register your client first",
						IsError: true,
					},
				},
			}
			stream.Send(out)
			continue
		}

		c := r.GetCode()
		g, ok := s.games[c]
		if !ok {
			out := &endless.Output{
				Data: &endless.Output_Msg{
					Msg: &endless.EventMessage{
						Msg:     fmt.Sprintf("No game with code '%v'", c),
						IsError: true,
					},
				},
			}
			stream.Send(out)
			continue
		}
		game = g
		break
	}

	return game.registerClient(stream)
}
