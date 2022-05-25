//
// SPDX-License-Identifier: Apache-2.0 OR MIT
//
// Copyright (C) 2022 Shun Sakai
//

// Package core provides the core of the program.
package core

import (
	"flag"
	"fmt"

	"github.com/sorairolake/jsonfmt/internal/cli"
	"github.com/sorairolake/jsonfmt/internal/info"
)

// Exit status.
const (
	// ExitSuccess is an exit status when program execution is successful.
	ExitSuccess = iota

	// ExitFailure is an exit status when an error occurred.
	ExitFailure
)

// Run returns the exit status of the program.
func Run() int {
	flag.Parse()

	var (
		args       = cli.Args
		inputFiles = flag.Args()
	)

	_ = inputFiles

	if args.Version {
		fmt.Fprintln(flag.CommandLine.Output(), info.CommandName, info.CommandVersion)

		return ExitSuccess
	}

	return ExitSuccess
}
