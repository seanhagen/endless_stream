package game

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/seanhagen/endless_stream/backend/endless"
)

// RegisterClient ...
func (g *Game) RegisterClient(id, name string, clientType endless.ClientType, stream endless.Game_StateServer) error {
	g.Lock()
	g.connected++
	log.Printf("client connected: %v, registering client", id)

	out, err := g.registerHuman(id, name, clientType)
	if err != nil {
		return err
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		cancel()
		g.unregisterHuman(out)
		//g.closingClients <- out
	}()

	isPlayer := out.isPlayer
	outCh := out.out

	g.Unlock()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			finished := false
			select {
			case _ = <-stream.Context().Done():
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
			case out := <-outCh:
				if stream.Context().Err() != nil {
					finished = true
					break
				}

				err := stream.Send(out)
				if err != nil {
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

// registerHuman is called above in RegisterClient
func (g *Game) registerHuman(id, name string, t endless.ClientType) (output, error) {
	// accessing some maps, gotta lock
	out := output{
		id:       id,
		out:      make(chan *endless.Output, 10),
		isPlayer: false,
	}

	v, ok := g.playerIds[id]
	if ok && v < 1 {
		// player is rejoining
		name = g.playerNames[id]
	}

	var err error
	var msg *endless.Output

	switch t {
	case endless.ClientType_ClientPlayer:
		if !g.started && len(g.players)+1 <= 4 {
			out, msg, err = g.registerPlayer(id, name, out)
		} else {
			out, msg, err = g.registerAudience(id, out)
		}
	case endless.ClientType_ClientDisplay:
		out, msg, err = g.registerDisplay(id, out)
	default:
		out, msg, err = g.registerAudience(id, out)
	}
	if msg != nil {
		time.AfterFunc(time.Millisecond*100, func() {
			g.output <- msg
		})
	}
	return out, err
}

// registerDisplay ...
func (g *Game) registerDisplay(id string, out output) (output, *endless.Output, error) {
	out.isDisplay = true
	g.displayClients[out] = true
	return out, nil, nil
}

// registerPlayer ...
func (g *Game) registerPlayer(id, name string, out output) (output, *endless.Output, error) {
	g.playerIds[id] = 1
	g.playerCharacters[id] = nil
	g.playerNames[id] = name

	isVip := false
	log.Printf("players connected: %v, first player? %v -- is vip? %v", len(g.players), len(g.players) == 0, id == g.vipPlayer)
	if len(g.players) == 0 || id == g.vipPlayer {
		isVip = true
		g.vipPlayer = id
	}

	log.Printf("registered player, id: %v, vip: %v", id, isVip)

	msg := &endless.Output{
		Data: &endless.Output_Joined{
			Joined: &endless.JoinedGame{
				Id:         id,
				AsAudience: false,
				IsVip:      isVip,
				Name:       name,
			},
		},
	}

	out.isPlayer = true
	g.players[out] = true

	return out, msg, nil
}

// registerAudience ...
func (g *Game) registerAudience(id string, out output) (output, *endless.Output, error) {
	msg := &endless.Output{
		Data: &endless.Output_Joined{
			Joined: &endless.JoinedGame{
				Id:         id,
				AsAudience: true,
				Name:       "Audience Member",
			},
		},
	}

	out.isAudience = true
	g.audience[out] = true

	return out, msg, nil
}
