package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/hashicorp/go-envparse"
	"github.com/magefile/mage/sh"
	"github.com/magefile/mage/target"
	"github.com/pterm/pterm"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	// can return an error for other reasons, but if this code can't
	// read the file then the Go cover tool probably can't either!
	return errors.Is(err, os.ErrNotExist)
}

func removeFile(path string) error {
	fmt.Printf("\tRemoving file %q: ", path)
	if !fileExists(path) {
		fmt.Printf(" file not found, skipping\n")
		return nil
	}

	err := os.Remove(path)
	if err != nil {
		fmt.Printf("%s - %s\n", color.RedString("ERROR"), err)
		return err
	}

	fmt.Printf("%s\n", color.GreenString("DONE"))

	return nil
}

func removeAll(path string) error {
	fmt.Printf("\tRemoving file or directory %q: ", path)
	if !fileExists(path) {
		fmt.Printf(" file or directory not found, skipping\n")
		return nil
	}

	err := os.RemoveAll(path)
	if err != nil {
		fmt.Printf("%s\n", color.RedString("ERROR"))
		return err
	}

	fmt.Printf("%s\n", color.GreenString("DONE"))
	return nil
}

func mkdir(path string) error {
	fmt.Printf("Creating directory %q...", path)
	if err := os.MkdirAll(path, 0o755); err != nil {
		fmt.Printf("%s - %s\n", color.RedString("ERROR"), err)
		return err
	}
	fmt.Printf("%s\n", color.GreenString("DONE"))
	return nil
}

func runCommand(cmd *exec.Cmd, area *pterm.AreaPrinter, action, errMsg string) error {
	buf := bytes.NewBuffer(nil)
	errBuf := bytes.NewBuffer(nil)
	cmd.Stdout = buf
	cmd.Stderr = errBuf

	if err := cmd.Run(); err != nil {
		addToArea(area, fmt.Sprintf("Unable to %s\n", action))
		addToArea(area, fmt.Sprintf("Output:\n%s\n", buf.String()))

		if errBuf.String() != "" {
			addToArea(area, fmt.Sprintf("Error:\n%s\n", errBuf.String()))
			return err
		}

		return errors.New(errMsg)
	}

	return area.Stop()
}

func addToArea(area *pterm.AreaPrinter, str string) {
	old := area.GetContent()
	new := fmt.Sprintf("%s\n%s", old, str)
	area.Update(new)
}

func getEnvVars() (map[string]string, error) {
	envFile, err := os.Open(".env")
	if err != nil {
		return nil, fmt.Errorf("unable to open '.env' file: %w", err)
	}

	envVars, err := envparse.Parse(envFile)
	if err != nil {
		return nil, fmt.Errorf("unable to get environment varaibles from .env: %w", err)
	}

	return envVars, nil
}

// outputsNewerThanInputs returns true if any file or directory from
// sources newer than a file or directory in destinations. It checks
// using [target.Dir]; each file or directory from destinations is
// passed as the first argument, sources is always passed as the
// second.
//
// For example, to check to see if any of the files generated from the GRPC protobuf schema files are older than the schema files themselves ( ie, the schema files have been updated )
//
// [target.Dir]:https://pkg.go.dev/github.com/magefile/mage@v1.15.0/target#Dir
func outputsNewerThanInputs(destinations, sources []string) (bool, error) {
	for _, dst := range destinations {
		sourceNewer, err := target.Dir(dst, sources...)
		if err != nil {
			return false, fmt.Errorf("can't check if sources newer than %q: %w", dst, err)
		}
		if sourceNewer {
			return true, nil
		}
	}

	return false, nil
}

func execCommand(command string, args ...string) error {
	outBuf := bytes.NewBuffer(nil)
	errBuf := bytes.NewBuffer(nil)

	ran, err := sh.Exec(nil, outBuf, errBuf, command, args...)
	if ran == false && err != nil {
		pterm.Error.Printf("Unable to run command!\n", command)
		fmt.Printf("Error:\n%s\n", errBuf.String())
		return fmt.Errorf("unable to run linter: %w", err)
	}

	if ran == true && err != nil {
		pterm.Warning.Printf("Errors while running %s!\n", command)
		pterm.Printf("Output:\n%s\n", outBuf.String())
	}

	if ran == true && err == nil {
		pterm.Success.Printf("No errors when running %s!\n", command)
	}

	return nil
}
