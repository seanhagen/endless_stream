package main

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/davecgh/go-spew/spew"
	"github.com/seanhagen/endless_stream/backend/endless"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	c, err := setupConn("localhost:10000")
	if err != nil {
		log.Fatal(err)
	}

	ec := endless.NewGameClient(c)

	resp, err := ec.Create(ctx, &endless.CreateGame{})
	if err != nil {
		log.Printf("unable to create game: %v", err)
		os.Exit(1)
	}

	strm, err := ec.State(ctx)
	if err != nil {
		log.Printf("error connecting to game state: %v", err)
		os.Exit(1)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			msg, err := strm.Recv()
			if err != nil {
				log.Printf("error recieved: %v", err)
				break
			} else {
				log.Printf("recieved message from game: %v", spew.Sdump(msg))
			}
		}
		wg.Done()
	}()

	err = strm.Send(&endless.Input{
		Input: &endless.Input_Register{
			Register: &endless.Register{
				Code: resp.GetCode(),
				Name: "client 1",
			},
		},
	})

	if err != nil {
		log.Printf("Unable to send message: %v", err)
	}

	wg.Wait()
	log.Printf("all done")
}

func setupConn(addr string) (*grpc.ClientConn, error) {
	dopts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithUserAgent("test client"),
	}
	return grpc.Dial(addr, dopts...)
}
