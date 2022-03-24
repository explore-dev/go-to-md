// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	gotomd "github.com/explore-dev/go-to-md"
)

var (
	// source file to process
	sourceFile = flag.String("file", "", "source filepath")

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

func commentsToString(filePath string, comments []string) string {
	var buf bytes.Buffer
	fileName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))

	sep := "\n"
	buf.WriteString(fmt.Sprintf("# %s", fileName))
	buf.WriteString(sep)
	buf.WriteString(strings.Join(comments, sep))
	return buf.String()
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

	fmt.Println(commentsToString(*sourceFile, gotomd.Collect(*sourceFile, *prefix, *annotation)))
}
