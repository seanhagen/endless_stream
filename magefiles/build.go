package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/pterm/pterm"
)

var (
	targetDirectory  = "target"
	serverBinaryName = "endless"
)

func Setup() error {
	if err := preBuildSetup(); err != nil {
		return fmt.Errorf("setup failed: %w", err)
	}

	if err := setupTargetDirectory(); err != nil {
		return fmt.Errorf("setup failed: %w", err)
	}

	return nil
}

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.SerialDeps(Clean, InstallDeps, Setup)
	fmt.Println("Building...")

	ldFlagsBase, err := buildLdFlags()
	if err != nil {
		return fmt.Errorf("unable to generate ldflags: %w", err)
	}

	args := []string{
		"build",
		"-o", serverBinaryPath(),
		"-ldflags=" + strings.Join(ldFlagsBase, " "),
		"./cmd",
	}

	area := pterm.DefaultArea.WithRemoveWhenDone(true)
	cmd := exec.Command(mg.GoCmd(), args...)
	return runCommand(cmd, area, "build the server binary", "failed to build server binary")
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Print("Installing dependencies...")

	goCmds := []string{
		"golang.org/x/vuln/cmd/govulncheck@latest",
	}

	for _, gc := range goCmds {
		if err := sh.Run("go", "install", gc); err != nil {
			pterm.Error.Printf("Failed to install %s\n", gc)
			return err
		}
	}

	fmt.Printf("%s\n", color.GreenString("Dependencies installed"))
	return nil
}

// Clean up after yourself
func Clean() error {
	fmt.Println("Cleaning up...")

	if err := removeAll(targetDirectory); err != nil {
		fmt.Printf("Cleaning up: %s", color.RedString("ERROR"))
		return err
	}

	fmt.Printf("Cleaning up: %s\n", color.GreenString("DONE"))
	return nil
}

func serverBinaryPath() string {
	return filepath.Join(targetDirectory, serverBinaryName)
}

// handles setting things up before a build happens
func preBuildSetup() error {
	return nil
}

// ensures target directory exists
func setupTargetDirectory() error {
	return mkdir(targetDirectory)
}

func buildLdFlags() ([]string, error) {
	version, err := sh.Output("cat", "VERSION")
	if err != nil {
		return nil, fmt.Errorf("unable to get version from VERSION file: %w", err)
	}

	build, err := sh.Output("git", "rev-parse", "HEAD")
	if err != nil {
		return nil, fmt.Errorf("unable to get git revision: %w", err)
	}

	flags := []string{
		"-X",
		fmt.Sprintf("\"main.Version=%s\"", version),
		"-X",
		fmt.Sprintf("\"main.Build=%s\"", build),
	}

	envVars, err := getEnvVars()
	if err != nil {
		return nil, fmt.Errorf("unable to get environment variables: %w", err)
	}

	for key, value := range envVars {
		addEnvVarToLdFlags(key, value, &flags)
	}

	flags = append(flags, "-s", "-w")

	return flags, nil
}

var _envToLdFlag = map[string]string{
	"REPO": "main.Repo",
}

func addEnvVarToLdFlags(key, value string, flags *[]string) {
	if !wantEnvVarAsLdFlag(key) {
		return
	}

	flagKey := _envToLdFlag[key]
	*flags = append(*flags,
		"-X", fmt.Sprintf("\"%s=%s\"", flagKey, value),
	)
}

func wantEnvVarAsLdFlag(key string) bool {
	_, exists := _envToLdFlag[key]
	return exists
}
