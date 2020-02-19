package game

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/seanhagen/endless_stream/backend/endless"
)

// RegisterClient ...
func (g *Game) RegisterClient(id, name string, stream endless.Game_StateServer) error {
	g.lock.Lock()
	defer g.lock.Unlock()
	log.Printf("client connected: %v, registering client", id)

	out, err := g.registerHuman(id, name)
	if err != nil {
		return err
	}
	log.Printf("human registered")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		cancel()
		g.closingClients <- out
	}()

	isPlayer := out.isPlayer
	outCh := out.out

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		log.Printf("client gofunc started")
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
		log.Printf("client incoming message gofunc started")
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
		log.Printf("client outgoing message gofunc started")
		// handle outgoing messages
		for {
			finished := false
			select {
			case <-ctx.Done():
				finished = true
			case out := <-outCh:
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

	log.Printf("waiting for client to finish")
	wg.Wait()
	log.Printf("client finished")
	return nil
}

// registerHuman is called above in RegisterClient
func (g *Game) registerHuman(id, name string) (output, error) {
	// accessing some maps, gotta lock
	out := output{
		id:       id,
		out:      make(chan *endless.Output),
		isPlayer: false,
	}

	v, ok := g.playerIds[id]
	if ok && v < 1 {
		// player is rejoining
		name = g.playerNames[id]
	}

	if len(g.players)+1 <= 4 {
		out.isPlayer = true
		log.Printf("client is player")
		g.newClients <- out
		log.Printf("output sent to channel")
		msg, err := g.registerPlayer(id, name)
		log.Printf("player registered")
		time.AfterFunc(time.Millisecond*100, func() {
			g.output <- msg
		})
		log.Printf("afterfunc registered")
		return out, err
	}

	g.newClients <- out
	msg, err := g.registerAudience(id)
	time.AfterFunc(time.Millisecond*100, func() {
		g.output <- msg
	})
	return out, err
}

// registerPlayer ...
func (g *Game) registerPlayer(id, name string) (*endless.Output, error) {
	g.playerIds[id] = 1
	g.playerCharacters[id] = nil
	g.playerNames[id] = name

	isVip := false
	if len(g.players) == 0 || id == g.vipPlayer {
		isVip = true
		g.vipPlayer = id
	}

	out := &endless.Output{
		Data: &endless.Output_Joined{
			Joined: &endless.JoinedGame{
				Id:         id,
				AsAudience: false,
				IsVip:      isVip,
				Name:       name,
			},
		},
	}

	return out, nil
}

// registerAudience ...
func (g *Game) registerAudience(id string) (*endless.Output, error) {
	out := &endless.Output{
		Data: &endless.Output_Joined{
			Joined: &endless.JoinedGame{
				Id:         id,
				AsAudience: true,
				Name:       "Audience Member",
			},
		},
	}

	return out, nil
}
