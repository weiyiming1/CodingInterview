package main

import (
	"CodingInterview/getOffer-Go/lib"
	"testing"
)

func TestPrintFromTopToBottom(t *testing.T) {
	d := lib.TreeNode{5, nil, nil}
	e := lib.TreeNode{7, nil, nil}
	f := lib.TreeNode{9, nil, nil}
	g := lib.TreeNode{11, nil, nil}
	b := lib.TreeNode{6, &d, &e}
	c := lib.TreeNode{10, &f, &g}
	a := lib.TreeNode{8, &b, &c}
	PrintFromTopToBottom(&a)
}
