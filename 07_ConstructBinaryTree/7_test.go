package main

import (
	"fmt"
	"testing"
)

func TestReConstructTree(t *testing.T) {
	preOrder := []int{1, 2, 4, 7}
	inOrder := []int{4, 7, 2, 1}
	root := ReConstructTree(preOrder, inOrder)
	fmt.Println(root)
}
