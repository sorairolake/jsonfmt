//
// SPDX-License-Identifier: Apache-2.0 OR MIT
//
// Copyright (C) 2022 Shun Sakai
//

// Package core provides the core of the program.
package core

import (
	"errors"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/sean-/sysexits"
	flag "github.com/spf13/pflag"

	"github.com/sorairolake/jsonfmt/internal/cli"
	"github.com/sorairolake/jsonfmt/internal/info"
)

// Run returns the exit status of the program.
func Run() int {
	flag.Parse()

	var (
		args       = cli.Args
		inputPaths = flag.Args()
	)

	if args.Version {
		fmt.Fprintln(os.Stderr, info.CommandName, info.CommandVersion)

		return ExitSuccess
	}

	inputFiles, err := readFiles(inputPaths)
	if err != nil {
		log.Err(err.(*os.PathError).Err).Msgf("Failed to read the file: %v:", err.(*os.PathError).Path)

		switch {
		case errors.Is(err, os.ErrNotExist):
			return sysexits.NoInput
		case errors.Is(err, os.ErrPermission):
			return sysexits.NoPerm
		default:
			return ExitFailure
		}
	}

	var outputFiles map[string][]byte

	if args.Compact {
		if output, err := batchCompact(inputFiles); err == nil {
			outputFiles = output
		} else {
			log.Err(err).Msg("Failed to format JSON:")

			return sysexits.DataErr
		}
	} else {
		indentLevel := int(args.Indent)
		if output, err := batchIndent(inputFiles, args.Tab, indentLevel); err == nil {
			outputFiles = output
		} else {
			log.Err(err).Msg("Failed to format JSON:")

			return sysexits.DataErr
		}
	}

	if args.Write {
		if err := writeFiles(outputFiles); err != nil {
			log.Err(err.(*os.PathError).Err).Msgf("Failed to write to the file: %v:", err.(*os.PathError).Path)

			switch {
			case errors.Is(err, os.ErrInvalid):
				return sysexits.CantCreate
			case errors.Is(err, os.ErrPermission):
				return sysexits.NoPerm
			default:
				return ExitFailure
			}
		}
	} else {
		for _, data := range outputFiles {
			fmt.Print(string(data))
		}
	}

	return ExitSuccess
}
