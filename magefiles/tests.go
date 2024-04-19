package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
	"github.com/pterm/pterm"
)

const (
	coverageProfile = "coverage.out"
	coverageOutput  = "coverage.txt"
)

func Tests() error {
	mg.SerialDeps(InstallDeps)

	if err := runTests(false); err != nil {
		pterm.Error.Printf("Failed: %s", err)
		os.Exit(1)
	}
	pterm.Success.Println("Tests run completed successfully.")

	return nil
}

// Coverage runs the unit tests and collects the code coverage data
func Coverage() error {
	mg.SerialDeps(InstallDeps)

	if err := runTests(true); err != nil {
		pterm.Error.Printf("Failed: %s", err)
		os.Exit(1)
	}
	pterm.Success.Println("Tests run completed successfully.")

	pterm.Info.Println("Generating code coverage.")
	if err := generateCoverage(); err != nil {
		pterm.Error.Println("Failed to generate code coverage data")
		return err
	}

	pterm.Success.Println("Code coverage generated.")
	return nil
}

func runTests(withCoverage bool) error {
	cmd := exec.Command("go", "test", "-v", "./...")

	area, err := pterm.DefaultArea.WithRemoveWhenDone().Start()
	if err != nil {
		log.Fatalf("unable to create pterm area: %s", err)
	}

	if withCoverage {
		addToArea(area, pterm.Info.Sprint("Running tests with coverage"))
		cmd.Args = append(cmd.Args, "-covermode=count", coverProfileArg())
	} else {
		addToArea(area, pterm.Info.Sprint("Running tests"))
	}
	addToArea(area, "\n")

	return runCommand(cmd, area, "run tests", "some tests failed to pass")
}

func generateCoverage() error {
	area, err := pterm.DefaultArea.WithRemoveWhenDone().Start()
	if err != nil {
		log.Fatalf("unable to create pterm area: %s", err)
	}

	area.Update(pterm.Info.Sprintf("Generating code coverage output "))
	if !fileExists(coverageProfile) {
		area.Update(pterm.Error.Sprint("ERROR"))
		return fmt.Errorf("unable to stat coverage file %q", coverageProfile)
	}

	cmd := exec.Command("go", "tool", "cover", goToolCoverFunc(), goToolCoverOutput()) //#nosec G204

	return runCommand(
		cmd,
		area,
		"generate code coverage data",
		"unable to generate code coverage data",
	)
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
