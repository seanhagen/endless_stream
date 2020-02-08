package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/davecgh/go-spew/spew"
	"github.com/seanhagen/endless_stream/backend/endless"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	c, err := setupConn("localhost:10001")
	if err != nil {
		log.Fatal(err)
	}

	ec := endless.NewGameClient(c)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("What would you like to do?\n\n\t1) Create new game\n\t2) Join game\n\t3) Re-join game\n\nEnter number:")
		txt, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Unable to read input: %v\n", err)
			os.Exit(1)
		}

		txt = strings.TrimSpace(txt)
		ch, err := strconv.Atoi(txt)
		if err != nil {
			fmt.Printf("\n'%v' is not a number.\n\n", txt)
			continue
		}

		switch ch {
		case 1:
			resp, err := ec.Create(ctx, &endless.CreateGame{})
			if err != nil {
				fmt.Printf("Unable to create game: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Game created, code: %v\nJoining game!\n\n", resp.GetCode())
			handleStreamInput(ctx, ec, resp.GetCode())
			goto complete
		case 2:
			fmt.Printf("\nEnter game code:")
			txt, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("\nUnable to read input: %v\n\n", err)
				os.Exit(1)
			}
			txt = strings.TrimSpace(txt)
			handleStreamInput(ctx, ec, txt)
			goto complete
		case 3:
			fmt.Printf("\nEnter game code:")
			code, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("\nUnable to read input: %v\n\n", err)
				os.Exit(1)
			}
			code = strings.TrimSpace(code)

			id, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("\nUnable to read input: %v\n\n", err)
				os.Exit(1)
			}
			id = strings.TrimSpace(id)

		default:
			fmt.Printf("Please enter 1 or 2.\n\n")
			continue
		}
	}
complete:
}

func handleStreamInput(ctx context.Context, client endless.GameClient, code string) {
	strm, err := client.State(ctx)
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
				Code: code,
				Name: "client 1",
			},
		},
	})

	if err != nil {
		log.Printf("Unable to send message: %v", err)
	}
	wg.Wait()
	fmt.Printf("Game done.\n\n")
}

func setupConn(addr string) (*grpc.ClientConn, error) {
	dopts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithUserAgent("test client"),
	}
	return grpc.Dial(addr, dopts...)
}
