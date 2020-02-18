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

	c, err := setupConn("localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	ec := endless.NewGameClient(c)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("What would you like to do?\n\n\t1) Create new game\n\t2) Join game\n\t3) Re-join game\n\t4) Join as audience\n\nEnter number:")
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

			fmt.Printf("What's your name? (Blank for anon): ")
			txt, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Unable to read input: %v\n", err)
				os.Exit(1)
			}
			name := strings.TrimSpace(txt)

			fmt.Printf("Game created, code: %v\nJoining game!\n\n", resp.GetCode())
			handleStreamInput(ctx, ec, resp.GetCode(), "", name)
			goto complete
		case 2:
			fmt.Printf("\nEnter game code:")
			txt, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("\nUnable to read input: %v\n\n", err)
				os.Exit(1)
			}
			code := strings.TrimSpace(txt)

			fmt.Printf("What's your name? (Blank for anon): ")
			txt, err = reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Unable to read input: %v\n", err)
				os.Exit(1)
			}
			name := strings.TrimSpace(txt)

			handleStreamInput(ctx, ec, code, "", name)
			goto complete
		case 3:
			fmt.Printf("\nEnter game code:")
			code, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("\nUnable to read input: %v\n\n", err)
				os.Exit(1)
			}
			code = strings.TrimSpace(code)

			fmt.Printf("\nEnter Player ID:")
			id, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("\nUnable to read input: %v\n\n", err)
				os.Exit(1)
			}
			id = strings.TrimSpace(id)

			handleStreamInput(ctx, ec, code, id, "")
			goto complete

		case 4:
			fmt.Printf("need audience code")
			goto complete
		default:
			fmt.Printf("Please enter 1 or 2.\n\n")
			continue
		}
	}
complete:
}

func handleStreamInput(ctx context.Context, client endless.GameClient, code, id, name string) {
	strm, err := client.State(ctx)
	if err != nil {
		log.Printf("error connecting to game state: %v", err)
		os.Exit(1)
	}

	l := &sync.Mutex{}
	var disp endless.Display

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			msg, err := strm.Recv()
			if err != nil {
				log.Printf("Error recieved: %v", err)
				break
			} else {
				switch t := msg.GetData().(type) {
				case *endless.Output_Tick:
					//ti := msg.GetTick()

				case *endless.Output_Selected:
					fmt.Printf("Character selected:\n")
					for k, v := range t.Selected.GetSelected() {
						fmt.Printf("\t%v -> %v\n", k, endless.ClassType_name[int32(v)])
					}

				case *endless.Output_State:
					s := msg.GetState()
					l.Lock()
					disp = s.GetDisplay()
					l.Unlock()

				case *endless.Output_Msg:
					m := msg.GetMsg()
					o := fmt.Sprintf("MSG[%v] ", m.GetMsgId())

					if pid := m.GetPlayerId(); pid != nil {
						o = fmt.Sprintf("%v Player[%v] ", o, pid.GetValue())
					}

					if m.GetIsError() {
						o = fmt.Sprintf("%v Error occured: %v", o, m.GetMsg())
					} else {
						if m.GetLogOnly() {
							o = fmt.Sprintf("%v LogOnly ", o)
						}

						if m.GetIsAlert() {
							o = fmt.Sprintf("%v ALERT ", o)
						}

						o = fmt.Sprintf("%v Message: %v", o, m.GetMsg())
					}

					log.Print(o)

				case *endless.Output_Joined:
					j := msg.GetJoined()
					log.Printf(
						"Player joined game:\n\tPlayer ID: %v -- %v\n\tVIP? %v\n\tAudience? %v",
						j.GetId(), j.GetName(), j.GetIsVip(), j.GetAsAudience())

				default:
					log.Printf("Unknown message type: %T", t)
				}
			}
		}
		wg.Done()
	}()

	err = strm.Send(&endless.Input{
		Input: &endless.Input_Register{
			Register: &endless.Register{
				Code: code,
				Name: name,
				Id:   id,
			},
		},
	})

	go func() {
		inp := bufio.NewReader(os.Stdin)
		for {
			switch disp {
			case endless.Display_ScreenCharSelect:
				fmt.Printf("Choose your character:\n\t1) Fighter\n\t2)Ranger\n\t3)Cleric\n\t4)Wizard\n\nEnter your selection [1-4]:")
				i := getIntFromInput(inp)
				var c endless.ClassType
				switch i {
				case 1:
					c = endless.ClassType_Fighter
				case 2:
					c = endless.ClassType_Ranger
				case 3:
					c = endless.ClassType_Cleric
				case 4:
					c = endless.ClassType_Wizard
				default:
					fmt.Printf("Not a valid selection, go again\n")
				}
				if c != endless.ClassType_Unknown {
					strm.Send(&endless.Input{
						Input: &endless.Input_CharSelect{
							CharSelect: &endless.CharSelect{
								PlayerId: id,
								Choice: &endless.Class{
									Class: c,
								},
							},
						},
					})
				}
			}
		}
	}()

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

type strRead interface {
	ReadString(byte) (string, error)
}

func getIntFromInput(r strRead) int {
	t, err := r.ReadString('\n')
	spew.Dump(t, err)
	if err != nil {
		return 0
	}
	t = strings.TrimSpace(t)
	ch, _ := strconv.Atoi(t)
	return ch
}
