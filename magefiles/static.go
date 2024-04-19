package main

import "github.com/magefile/mage/mg"

func Lint() error {
	mg.SerialDeps(InstallDeps)

	args := []string{
		"run", "-c", "./.golangci.yml",
		"--out-format=colored-line-number",
		"--color", "always",
	}

	return execCommand("golangci-lint", args...)
}

func VulnerabilityScan() error {
	mg.SerialDeps(InstallDeps)

	return execCommand("govulncheck", "./...")
}
