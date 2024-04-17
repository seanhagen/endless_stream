//go:build mage
// +build mage

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

func GRPC() error {
	return nil
}

// Coverage runs the unit tests and collects the code coverage data
func Coverage() error {
	mg.Deps(InstallDeps)
	fmt.Println("Running tests...")

	cmd := exec.Command(
		"go",
		"test",
		"-v",
		"./...",
		"-covermode=count",
		"-coverprofile=coverage.out",
	)

	buf := bytes.NewBuffer(nil)
	cmd.Stdout = buf

	if err := cmd.Run(); err != nil {
		fmt.Printf("Unable to run tests: %s\n", err)
		fmt.Printf("Error:\n%s\n", buf.String())
		return err
	}

	cmd = exec.Command("go", "tool", "cover", "-func=coverage.out", "-o=coverage.txt")
	buf.Reset()
	cmd.Stdout = buf

	if err := cmd.Run(); err != nil {
		fmt.Printf("Unable to generate coverage data: %s\n", err)
		fmt.Printf("Error:\n%s\n", buf.String())
		return err
	}

	return nil
}

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

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
	cmd := exec.Command("go", "get", "github.com/stretchr/piglatin")
	return cmd.Run()
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("MyApp")
}
