package game

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/seanhagen/endless_stream/backend/grpc"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

type Srv struct {
	games map[string]*Game
}

func getGameId() string {
	b := make([]byte, 4)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// Create ...
func (s *Srv) Create(ctx context.Context, in *grpc.CreateGame) (*grpc.GameCreated, error) {
	var id string
	for {
		id = getGameId()
		if _, ok := s.games[id]; ok {
			continue
		}

		g, err := createGame(id)
		if err != nil {
			return nil, err
		}
		s.games[id] = g
		break
	}

	return &grpc.GameCreated{Code: id}, nil
}

// State ...
func (s *Srv) State(stream grpc.GameServer_StateServer) error {
	var game *Game
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Printf("Unable to receive state message during handshake: %v", err)
			return fmt.Errorf("unable to receive message? %v", err)
		}

		r := msg.GetRegister()
		if r == nil {
			out := &grpc.Output{
				Data: &grpc.Output_Msg{
					Msg: &grpc.EventMessage{
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
			out := &grpc.Output{
				Data: &grpc.Output_Msg{
					Msg: &grpc.EventMessage{
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
