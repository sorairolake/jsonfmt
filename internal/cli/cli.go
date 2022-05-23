//
// SPDX-License-Identifier: Apache-2.0 OR MIT
//
// Copyright (C) 2022 Shun Sakai
//

package cli

import (
	"flag"
	"fmt"

	"github.com/sorairolake/jsonfmt/internal/info"
)

type Opt struct {
	Compact bool
	Tab     bool
	Indent  int
	Write   bool
	Version bool
}

var Args Opt

func init() {
	flag.BoolVar(&Args.Compact, "compact", false, "Print JSON on a single-line")
	flag.BoolVar(&Args.Tab, "tab", false, "Indent with tabs instead of spaces")
	flag.IntVar(&Args.Indent, "indent", 2, "Number of spaces per indentation level")
	flag.BoolVar(&Args.Write, "write", false, "Edit files in-place")
	flag.BoolVar(&Args.Version, "version", false, "Print version information")

	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "Usage:", info.CommandName, "[OPTIONS] [FILE]...")
		flag.PrintDefaults()
	}
}
