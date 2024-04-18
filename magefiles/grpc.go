package main

import (
	"fmt"
)

func GRPC() error {
	if err := runBuf(); err != nil {
		return err
	}

	return nil
}

// RunBuf ...
func runBuf() error {
	fmt.Printf("Running buf to generate GRPC code...\n")

	return nil
}
