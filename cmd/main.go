/*
Copyright © 2024 Sean Patrick Hagen <sean.hagen@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package main

import "os"

var (
	// Version is set by the build process, contains semantic version.
	Version string //nolint:gochecknoglobals
	// Build is set by the build process, contains sha tag of build.
	Build string //nolint:gochecknoglobals
	// Repo is set by the build process, contains the repo where the code for this binary was built from.
	Repo string //nolint:gochecknoglobals
)

func main() {
	rootCmd.AddCommand(serveCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
