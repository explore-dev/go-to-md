// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package test

// funcWithNoParamsNoReturn is a Go function.
// It has an empty signature and body.
func funcWithNoParamsNoReturn() { // comment at the header of the func
	// comment inside of the func
}

func funcWithParamsNoReturn(s string) {
}

/*


   revy-doc: test
   ## funcWithNoParamsReturn
*/
// funcWithNoParamsReturn is a Go function.
// It returns an empty string.
func funcWithNoParamsReturn() string {
	// This is a function
	kf := func() string {
		return ""
	}
	return kf()
}

//    revy-doc: test
/* ## funcWithParamsReturn
   This function receives some arguments and returns the empty string.
*/
func funcWithParamsReturn(i int, s string, rest ...string) string {
	return ""
}

/*
   revy-doc: ignoring
### funcWithMultipleParamsReturn
This comment on the function will be ignored.
*/
func funcWithMultipleParamsReturn(i, k int, s string, rest ...string) string {
	return ""
}

func funcWithMultipleParamsMultipleReturn(i, k int, s string, rest ...string) (string, error) {
	return "", nil
}

type Receiver struct{}

/*
revy-doc: test
### methodOfReceiverNoParamsNoReturn
Method on Receiver object that will do nothing.
*/
func (r *Receiver) methodOfReceiverNoParamsNoReturn() {

}

func (r *Receiver) methodOfReceiverNoParamsReturn() string {
	return ""
}
