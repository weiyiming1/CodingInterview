package main

import (
	"CodingInterview/getOffer-Go/lib"
	"log"
	"testing"
)

func TestHasSubTree(t *testing.T) {
	a := lib.TreeNode{4, nil, nil}
	b := lib.TreeNode{7, nil, nil}
	c := lib.TreeNode{2, &a, &b}
	d := lib.TreeNode{9, nil, nil}
	e := lib.TreeNode{8, &d, &c}
	f := lib.TreeNode{7, nil, nil}
	g := lib.TreeNode{8, &e, &f}

	aa := lib.TreeNode{9, nil, nil}
	bb := lib.TreeNode{2, nil, nil}
	cc := lib.TreeNode{8, &aa, &bb}

	res := HasSubTree(&g, &cc)
	log.Println(res)
}
