// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	gotomd "github.com/explore-dev/go-to-md"
)

var (
	// source file to process
	sourceFile = flag.String("file", "", "souce file to process")

	// comment prefix
	prefix = flag.String("prefix", "", "prefix to identify comment annotations")

	// annotation
	annotation = flag.String("annotation", "", "annotation that marks func declarations comments in markdown")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: go-to-md -file=LOC -prefix=PREFIX -annotation=ANNOTATION\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Parse()

	if flag.NArg() > 0 {
		fmt.Fprintln(os.Stderr, `Unexpected arguments.`)
		usage()
	}
	if *sourceFile == "" || *prefix == "" || *annotation == "" {
		fmt.Fprintln(os.Stderr, "At least one of -file, -prefix, or -annotation is a zero value.")
		usage()
	}

	name := strings.TrimSuffix(filepath.Base(*sourceFile), filepath.Ext(*sourceFile))
	fmt.Printf("# %s\n", name)
	for _, comment := range gotomd.Collect(*sourceFile, *prefix, *annotation) {
		fmt.Println(comment)
	}
}
