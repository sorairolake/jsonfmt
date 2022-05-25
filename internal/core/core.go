//
// SPDX-License-Identifier: Apache-2.0 OR MIT
//
// Copyright (C) 2022 Shun Sakai
//

package core

import (
	"flag"
	"fmt"

	"github.com/sorairolake/jsonfmt/internal/cli"
	"github.com/sorairolake/jsonfmt/internal/info"
)

const (
	exitSuccess = iota
	exitFailure
)

func Run() int {
	flag.Parse()
	args := cli.Args

	if args.Version {
		fmt.Fprintln(flag.CommandLine.Output(), info.CommandName, info.CommandVersion)

		return exitSuccess
	}

	return exitSuccess
}
