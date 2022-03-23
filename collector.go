// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package gotomd

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

func commentHasAnnotation(comment, prefix, annotation string) bool {
	normalizedComment := strings.ToLower(strings.Trim(comment, " "))
	normalizedAnnotation := strings.ToLower(annotation)
	if strings.HasPrefix(normalizedComment, prefix) {
		rest := normalizedComment[len(prefix)-1:]
		annotations := strings.Split(rest, " ")
		for _, annot := range annotations {
			if normalizedAnnotation == annot {
				return true
			}
		}
	}

	return false
}

func Collect(file, prefix, annotation string) []string {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	fnComments := make([]string, 0)

	for _, decl := range astFile.Decls {
		if fd, ok := decl.(*ast.FuncDecl); ok {
			if fd.Doc == nil || len(fd.Doc.List) == 0 {
				continue
			}

			docText := fd.Doc.Text()
			comments := strings.Split(docText, "\n")
			if len(comments) > 0 {
				if commentHasAnnotation(comments[0], prefix, annotation) {
					docTextWithoutAnnotation := strings.Join(comments[1:], "\n")
					fnComments = append(fnComments, docTextWithoutAnnotation)
				}
			}
		}
	}

	return fnComments
}
