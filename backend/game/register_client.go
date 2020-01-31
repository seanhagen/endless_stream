package game

import (
	"context"
	"log"
	"sync"

	"github.com/gofrs/uuid"
	"github.com/seanhagen/endless_stream/backend/endless"
)

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

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer func() {
		cancel()
		g.closingClients <- output
	}()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for {
			finished := false
			select {
			case _ = <-stream.Context().Done():
				log.Printf("Game %v stream client context done", g.code)
				finished = true
			case <-ctx.Done():
				finished = true
			default:
			}

			if finished {
				break
			}
		}

		cancel()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		// handle incoming messages
		for {
			finished := false
			select {
			case <-ctx.Done():
				finished = true
			default:
				// this is a blocking operation, so this will wait until either:
				//   A) a message is received from the client
				//   B) the client disconnects ( causing an error )
				//   C) something else happens that causes an error
				msg, err := stream.Recv()
				if err != nil {
					finished = true
					break
				}

				msg.PlayerId = id.String()
				g.input <- input{in: msg, isPlayer: isPlayer}
			}

			if finished {
				break
			}
		}

		cancel()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		// handle outgoing messages
		for {
			finished := false
			select {
			case <-ctx.Done():
				finished = true
			case out := <-output.out:
				if stream.Context().Err() != nil {
					finished = true
					break
				}

				err := stream.Send(out)
				if err != nil {
					log.Printf("Game %v unable to send message to client: %v", g.code, err)
					finished = true
					break
				}
			}
			if finished {
				break
			}
		}

		cancel()
		wg.Done()
	}()

	wg.Wait()
	return nil
}
