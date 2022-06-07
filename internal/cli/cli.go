//
// SPDX-License-Identifier: Apache-2.0 OR MIT
//
// Copyright (C) 2022 Shun Sakai
//

// Package cli provides the command-line interface.
package cli

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/sorairolake/jsonfmt/internal/info"
)

// Opt represents structure of the command-line interface.
type Opt struct {
	// Compact is an argument for printing JSON on a single-line.
	Compact bool

	// Tab is an argument for indenting with tabs instead of spaces.
	Tab bool

	// Indent is an argument for number of spaces per indentation level.
	Indent uint8

	// Write is an argument for editing files in-place.
	Write bool

	// Version is an argument for printing version information.
	Version bool
}

// Args is the structure of the command-line interface.
var Args Opt

func init() {
	flag.BoolVarP(&Args.Compact, "compact", "c", false, "Print JSON on a single-line")
	flag.BoolVar(&Args.Tab, "tab", false, "Indent with tabs instead of spaces")
	flag.Uint8Var(&Args.Indent, "indent", 2, "Number of spaces per indentation level")
	flag.BoolVarP(&Args.Write, "write", "w", false, "Edit files in-place")
	flag.BoolVarP(&Args.Version, "version", "V", false, "Print version information")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %v [OPTIONS] [FILE]...\n", info.CommandName)

		flag.PrintDefaults()
	}
}
