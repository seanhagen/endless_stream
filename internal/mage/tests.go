package mage

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

const (
	coverageProfile = "coverage.out"
	coverageOutput  = "coverage.txt"
)

func coverProfileArg() string {
	return fmt.Sprintf("-coverprofile=%s", coverageProfile)
}

func goToolCoverFunc() string {
	return fmt.Sprintf("-func=%s", coverageProfile)
}

func goToolCoverOutput() string {
	return fmt.Sprintf("-o=%s", coverageOutput)
}

func coverProfileExists() bool {
	_, err := os.Stat(coverageProfile)
	if err == nil {
		return true
	}

	// can return an error for other reasons, but if this code can't
	// read the file then the Go cover tool probably can't either!
	return errors.Is(err, os.ErrNotExist)
}

func RunTests(withCoverage bool) error {
	cmd := exec.Command("go", "test", "-v", "./...")

	fmt.Printf("Running tests ")
	if withCoverage {
		fmt.Printf("(with coverage)")
		cmd.Args = append(cmd.Args, "-covermode=count", coverProfileArg())
	}
	fmt.Printf("\n")

	buf := bytes.NewBuffer(nil)
	cmd.Stdout = buf

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Unable to run tests: %s\n", err)
		fmt.Printf("Error:\n%s\n", buf.String())
	}

	return err
}

func GenerateCoverage() error {
	fmt.Printf("Generating coverage output...\n")
	if !coverProfileExists() {
		return fmt.Errorf("unable to stat coverage file %q", coverageProfile)
	}

	cmd := exec.Command("go", "tool", "cover", goToolCoverFunc(), goToolCoverOutput())

	buf := bytes.NewBuffer(nil)
	cmd.Stdout = buf

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Unable to generate coverage data: %s\n", err)
		fmt.Printf("Error:\n%s\n", buf.String())
		return err
	}

	return err
}
