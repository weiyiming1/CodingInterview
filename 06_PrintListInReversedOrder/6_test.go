package main

import "testing"

func TestPrintLinkReverse(t *testing.T) {
	a := &LinkNode{
		5,
		nil,
	}
	b := &LinkNode{
		4,
		a,
	}

	c := &LinkNode{
		3,
		b,
	}

	PrintLinkReverse(c)
}
