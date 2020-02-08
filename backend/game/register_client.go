package game

import (
	"context"
	"log"
	"sync"

	"github.com/seanhagen/endless_stream/backend/endless"
)

// registerClient ...
func (g *Game) registerClient(id string, stream endless.Game_StateServer) error {
	msg, out, err := g.registerHuman(id)
	if err != nil {
		return err
	}
	stream.Send(msg)

	g.newClients <- out
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		cancel()
		g.closingClients <- out
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

				msg.PlayerId = id
				g.input <- input{in: msg, isPlayer: out.isPlayer}
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
			case out := <-out.out:
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
