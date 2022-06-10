//
// SPDX-License-Identifier: Apache-2.0 OR MIT
//
// Copyright (C) 2022 Shun Sakai
//

package core

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFilesFromStdin(t *testing.T) {
	t.Parallel()

	tempFile, err := os.CreateTemp("", "test-read-files-from-stdin-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write([]byte("Hello, world!")); err != nil {
		t.Fatal(err)
	}

	if _, err := tempFile.Seek(0, 0); err != nil {
		t.Fatal(err)
	}

	origStdin := os.Stdin
	defer func() { os.Stdin = origStdin }()
	os.Stdin = tempFile

	readFile, err := readFiles([]string{})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Hello, world!", string(readFile[""]))
}

func TestReadFilesFromFile(t *testing.T) {
	t.Parallel()

	tempDir, err := os.MkdirTemp("", "test-read-files-from-file-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	filePath := filepath.Join(tempDir, "file")
	if err := os.WriteFile(filePath, []byte("Hello, world!"), 0o644); err != nil {
		t.Fatal(err)
	}

	readFile, err := readFiles([]string{filePath})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Hello, world!", string(readFile[filePath]))
}

func TestBatchCompact(t *testing.T) {
	t.Parallel()

	tempDir, err := os.MkdirTemp("", "test-batch-compact-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	filePath := filepath.Join(tempDir, "file")
	const inputJSON = `{
  "foo": "Hello, world!",
  "bar": 255,
  "baz": null
}`
	input := map[string][]byte{filePath: []byte(inputJSON)}

	output, err := batchCompact(input)
	if err != nil {
		t.Fatal(err)
	}

	const expectedJSON = `{"foo":"Hello, world!","bar":255,"baz":null}` + "\n"
	assert.Equal(t, expectedJSON, string(output[filePath]))
}

func TestBatchIndentWithSpaces(t *testing.T) {
	t.Parallel()

	tempDir, err := os.MkdirTemp("", "test-batch-indent-with-spaces-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	filePath := filepath.Join(tempDir, "file")
	const inputJSON = `{"foo":"Hello, world!","bar":255,"baz":null}`
	input := map[string][]byte{filePath: []byte(inputJSON)}

	output, err := batchIndent(input, false, 2)
	if err != nil {
		t.Fatal(err)
	}

	const expectedJSON = `{
  "foo": "Hello, world!",
  "bar": 255,
  "baz": null
}`
	assert.Equal(t, expectedJSON, string(output[filePath]))
}

func TestBatchIndentWithTabs(t *testing.T) {
	t.Parallel()

	tempDir, err := os.MkdirTemp("", "test-batch-indent-with-tabs-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	filePath := filepath.Join(tempDir, "file")
	const inputJSON = `{"foo":"Hello, world!","bar":255,"baz":null}`
	input := map[string][]byte{filePath: []byte(inputJSON)}

	output, err := batchIndent(input, true, 1)
	if err != nil {
		t.Fatal(err)
	}

	const expectedJSON = `{
	"foo": "Hello, world!",
	"bar": 255,
	"baz": null
}`
	assert.Equal(t, expectedJSON, string(output[filePath]))
}

func TestWriteFiles(t *testing.T) {
	t.Parallel()

	tempDir, err := os.MkdirTemp("", "test-write-files-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	filePath := filepath.Join(tempDir, "file")
	input := map[string][]byte{filePath: []byte("Hello, world!")}
	if err := writeFiles(input); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Hello, world!", string(data))
}
