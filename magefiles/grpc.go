package main

import (
	"fmt"
	"os/exec"

	"github.com/pterm/pterm"
)

func GRPC() error {
	destinations := []string{
		"internal/proto",
		"pb",
		"docs/openapiv2",
		"docs/grpc",

		// the C# files have to be listed individually for raisins
		"EndlessStreamData/Hex.cs",
		"EndlessStreamData/HexService.cs",
		"EndlessStreamData/HexServiceGrpc.cs",
	}

	sources := []string{
		"proto",
	}

	rebuild, err := outputsNewerThanInputs(destinations, sources)
	if err != nil {
		return fmt.Errorf(
			"unable to determine if GRPC protobuf schemas are newer than the generated files: %w",
			err,
		)
	}

	if !rebuild {
		pterm.Info.Println(
			"GRPC protobuf schemas are older than all the generated code, not rebuilding.",
		)
		return nil
	}

	return runBuf()
}

// RunBuf ...
func runBuf() error {
	fmt.Printf("Running buf to generate GRPC code...\n")

	cmd := exec.Command("buf", "generate")
	area := pterm.DefaultArea.WithRemoveWhenDone(true)
	return runCommand(
		cmd,
		area,
		"generating GRPC from protobufs schemas",
		"failed to generate GRPC code from protobuf schemas",
	)
}
