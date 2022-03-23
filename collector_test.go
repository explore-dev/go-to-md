// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package gotomd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentHasAnnotationTrue(t *testing.T) {
	comment := "   prefix: test"
	annotation := "Test"

	got := commentHasAnnotation(comment, "prefix", annotation)

	assert.True(t, got)
}

func TestCommentHasAnnotationFalseAnnotation(t *testing.T) {
	comment := "   prefix: test"
	annotation := "Tes"

	got := commentHasAnnotation(comment, "prefix", annotation)

	assert.False(t, got)
}

func TestCommentHasAnnotationFalseComent(t *testing.T) {
	comment := "prfix: test"
	annotation := "test"

	got := commentHasAnnotation(comment, "prefix", annotation)

	assert.False(t, got)
}

func TestCollector(t *testing.T) {
	commentOnFuncWithNoParamsReturn :=
		`   ## funcWithNoParamsReturn

funcWithNoParamsReturn is a Go function.
It returns an empty string.
`
	commentOnFuncWithParamsReturn :=
		` ## funcWithParamsReturn
   This function receives some arguments and returns the empty string.
`

	commentOnMethodOfReceiverNoParamsNoReturn :=
		`### methodOfReceiverNoParamsNoReturn
Method on Receiver object that will do nothing.
`
	wantVal := []string{commentOnFuncWithNoParamsReturn, commentOnFuncWithParamsReturn, commentOnMethodOfReceiverNoParamsNoReturn}

	gotVal := Collect("./resources/test/foo.go", "revy-doc", "test")

	fmt.Println(gotVal)
	assert.Equal(t, wantVal, gotVal)
}
