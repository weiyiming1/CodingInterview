package main

import (
	"log"
	"testing"
)

func TestFindKthToTail(t *testing.T) {
	f := Node{6, nil}
	e := Node{5, &f}
	d := Node{4, &e}
	c := Node{3, &d}
	b := Node{2, &c}
	a := Node{1, &b}
	target := FindKthToTail(&a, 3)
	log.Println(target)
}
