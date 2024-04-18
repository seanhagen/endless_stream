//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/seanhagen/endless_stream/internal/mage"
)

func GRPC() error {
	mg.Deps(InstallDeps)

	if err := mage.RunBuf(); err != nil {
		return err
	}

	return nil
}

// Coverage runs the unit tests and collects the code coverage data
func Coverage() error {
	mg.Deps(InstallDeps)

	fmt.Printf("Generating code coverage...\n")

	if err := mage.RunTests(true); err != nil {
		return err
	}

	if err := mage.GenerateCoverage(); err != nil {
		return err
	}

	fmt.Printf("%s - finished generating code coverage\n", color.GreenString("DONE"))

	return nil
}

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(InstallDeps)
	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "-o", "MyApp", ".")
	return cmd.Run()
}

// A custom install step if you need your bin someplace other than go/bin
func Install() error {
	mg.Deps(Build)
	fmt.Println("Installing...")
	return os.Rename("./MyApp", "/usr/bin/MyApp")
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	// cmd := exec.Command("go", "get", "github.com/stretchr/piglatin")
	// return cmd.Run()
	return nil
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("MyApp")
}
