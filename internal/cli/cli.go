//
// SPDX-License-Identifier: Apache-2.0 OR MIT
//
// Copyright (C) 2022 Shun Sakai
//

// Package cli provides the command-line interface.
package cli

import (
	"flag"
	"fmt"

	"github.com/sorairolake/jsonfmt/internal/info"
)

// Opt represents structure of the command-line interface.
type Opt struct {
	// Compact is an argument for printing JSON on a single-line.
	Compact bool

	// Tab is an argument for indenting with tabs instead of spaces.
	Tab bool

	// Indent is an argument for number of spaces per indentation level.
	Indent uint

	// Write is an argument for editing files in-place.
	Write bool

	// Version is an argument for printing version information.
	Version bool
}

// Args is the structure of the command-line interface.
var Args Opt

func init() {
	flag.BoolVar(&Args.Compact, "compact", false, "Print JSON on a single-line")
	flag.BoolVar(&Args.Tab, "tab", false, "Indent with tabs instead of spaces")
	flag.UintVar(&Args.Indent, "indent", 2, "Number of spaces per indentation level")
	flag.BoolVar(&Args.Write, "write", false, "Edit files in-place")
	flag.BoolVar(&Args.Version, "version", false, "Print version information")

	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), info.CommandName, info.CommandVersion)
		fmt.Fprintf(flag.CommandLine.Output(), "A formatter for JSON\n\n")

		fmt.Fprintln(flag.CommandLine.Output(), "Usage:")
		fmt.Fprintf(flag.CommandLine.Output(), "  %v [OPTIONS] [FILE]...\n\n", info.CommandName)

		fmt.Fprintln(flag.CommandLine.Output(), "Args:")
		fmt.Fprintln(flag.CommandLine.Output(), "  <FILE>...")
		fmt.Fprintf(flag.CommandLine.Output(), "    \tFiles to format\n\n")

		fmt.Fprintln(flag.CommandLine.Output(), "Options:")
		flag.PrintDefaults()
		fmt.Fprintln(flag.CommandLine.Output(), "  -h, -help")
		fmt.Fprintln(flag.CommandLine.Output(), "    \tPrint help information")
	}
}
