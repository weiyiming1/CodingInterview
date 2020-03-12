package main

import (
	"CodingInterview/getOffer-Go/lib"
	"log"
	"testing"
)

func TestSerialize(t *testing.T) {
	t4 := lib.TreeNode{4, nil, nil}
	t5 := lib.TreeNode{5, nil, nil}
	t6 := lib.TreeNode{6, nil, nil}
	t2 := lib.TreeNode{2, &t4, nil}
	t3 := lib.TreeNode{3, &t5, &t6}
	t1 := lib.TreeNode{1, &t2, &t3}

	s := []string{}
	s = Serialize(&t1, s)
	var a1 *lib.TreeNode
	Deserialize(&a1, &s)
	log.Println("a1:", a1)
	a2 := a1.Left
	log.Println("a2:", a2)
	a3 := a1.Right
	log.Println("a3:", a3)
	a4 := a2.Left
	log.Println("a4:", a4)
	a5 := a3.Left
	log.Println("a5:", a5)
	a6 := a3.Right
	log.Println("a6:", a6)
}
