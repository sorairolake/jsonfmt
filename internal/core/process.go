//
// SPDX-License-Identifier: Apache-2.0 OR MIT
//
// Copyright (C) 2022 Shun Sakai
//

package core

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strings"
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
		if err := os.WriteFile(path, data, 0o644); err != nil {
			return err
		}
	}

	return nil
}
