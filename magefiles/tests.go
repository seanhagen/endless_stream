package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
)

const (
	coverageProfile = "coverage.out"
	coverageOutput  = "coverage.txt"
)

// Coverage runs the unit tests and collects the code coverage data
func Coverage() error {
	mg.Deps(InstallDeps)

	fmt.Printf("Generating code coverage...\n")

	if err := runTests(true); err != nil {
		return err
	}

	if err := generateCoverage(); err != nil {
		return err
	}

	fmt.Printf("%s - finished generating code coverage\n", color.GreenString("DONE"))

	return nil
}

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

func runTests(withCoverage bool) error {
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

func generateCoverage() error {
	fmt.Printf("Generating coverage output...\n")
	if !coverProfileExists() {
		return fmt.Errorf("unable to stat coverage file %q", coverageProfile)
	}

	cmd := exec.Command("go", "tool", "cover", goToolCoverFunc(), goToolCoverOutput()) //#nosec G204

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
