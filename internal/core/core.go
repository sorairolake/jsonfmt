//
// SPDX-License-Identifier: Apache-2.0 OR MIT
//
// Copyright (C) 2022 Shun Sakai
//

// Package core provides the core of the program.
package core

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

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

func readFiles(inputPaths []string) (map[string][]byte, error) {
	inputFiles := make(map[string][]byte)

	if len(inputPaths) == 0 {
		if data, err := io.ReadAll(os.Stdin); err == nil {
			inputFiles[""] = data
		} else {
			return nil, err
		}
	} else {
		for _, inputPath := range inputPaths {
			if data, err := os.ReadFile(inputPath); err == nil {
				inputFiles[inputPath] = data
			} else {
				return nil, err
			}
		}
	}

	return inputFiles, nil
}

func batchCompact(inputFiles map[string][]byte) (map[string][]byte, error) {
	for path, data := range inputFiles {
		var buf bytes.Buffer
		if err := json.Compact(&buf, data); err != nil {
			return nil, err
		}
		inputFiles[path] = buf.Bytes()
	}

	return inputFiles, nil
}

func batchIndent(inputFiles map[string][]byte, isUseTab bool, indentLevel int) (map[string][]byte, error) {
	var indent string
	if isUseTab {
		indent = "\t"
	} else {
		indent = " "
	}
	indent = strings.Repeat(indent, indentLevel)

	for path, data := range inputFiles {
		var buf bytes.Buffer
		if err := json.Indent(&buf, data, "", indent); err != nil {
			return nil, err
		}
		inputFiles[path] = buf.Bytes()
	}

	return inputFiles, nil
}

func writeFiles(inputFiles map[string][]byte) error {
	for path, data := range inputFiles {
		if err := os.WriteFile(path, data, 0644); err != nil {
			return err
		}
	}

	return nil
}

// Run returns the exit status of the program.
func Run() int {
	flag.Parse()

	var (
		args       = cli.Args
		inputPaths = flag.Args()
	)

	if args.Version {
		fmt.Fprintln(flag.CommandLine.Output(), info.CommandName, info.CommandVersion)

		return ExitSuccess
	}

	inputFiles, err := readFiles(inputPaths)
	if err != nil {
		log.Print(err)

		return ExitFailure
	}

	var outputFiles map[string][]byte

	if args.Compact {
		if output, err := batchCompact(inputFiles); err == nil {
			outputFiles = output
		} else {
			log.Print(err)

			return ExitFailure
		}
	} else {
		indentLevel := int(args.Indent)
		if output, err := batchIndent(inputFiles, args.Tab, indentLevel); err == nil {
			outputFiles = output
		} else {
			log.Print(err)

			return ExitFailure
		}
	}

	if args.Write {
		if err := writeFiles(outputFiles); err != nil {
			log.Print(err)

			return ExitFailure
		}
	} else {
		for _, data := range outputFiles {
			fmt.Println(string(data))
		}
	}

	return ExitSuccess
}
