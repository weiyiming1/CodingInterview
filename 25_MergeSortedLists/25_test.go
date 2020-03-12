package main

import (
	"log"
	"testing"
)

func TestMerge(t *testing.T) {
	d := Node{7, nil}
	c := Node{5, &d}
	b := Node{3, &c}
	a := Node{1, &b}

	dd := Node{8, nil}
	cc := Node{6, &dd}
	bb := Node{4, &cc}
	aa := Node{2, &bb}

	mergedHead := Merge(&a, &aa)
	for mergedHead != nil {
		log.Println(mergedHead.Val)
		mergedHead = mergedHead.Next
	}
}
