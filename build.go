//
// SPDX-License-Identifier: Apache-2.0 OR MIT
//
// Copyright (C) 2022 Shun Sakai
//

package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sorairolake/jsonfmt/internal/info"
)

func generateManPage(outDir string) error {
	var manDir string
	if dir, err := os.Getwd(); err == nil {
		manDir = filepath.Join(dir, "doc/man/man1")
	} else {
		return err
	}

	args := []string{"-b", "manpage", "-a", strings.Join([]string{"revnumber", info.CommandVersion}, "="), "-D", outDir, filepath.Join(manDir, "jsonfmt.1.adoc")}

	return exec.Command("asciidoctor", args...).Run()
}

func main() {
	var outDir string
	if dir, err := os.Getwd(); err == nil {
		outDir = filepath.Join(dir, "build")
	} else {
		log.Fatal(err)
	}

	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		if err := os.Mkdir(filepath.Base(outDir), 0755); err != nil {
			log.Fatal(err)
		}
	}

	if err := generateManPage(outDir); err != nil {
		log.Fatalf("Asciidoctor failed (%v)", err)
	}
}
