package main

import (
	"CodingInterview/getOffer-Go/lib"
	"log"
	"testing"
)

func TestIsBalanceTree(t *testing.T) {
	n7 := lib.TreeNode{7, nil, nil}
	n4 := lib.TreeNode{4, nil, nil}
	n5 := lib.TreeNode{5, &n7, nil}
	n2 := lib.TreeNode{2, &n4, &n5}
	n6 := lib.TreeNode{6, nil, nil}
	n3 := lib.TreeNode{3, nil, &n6}
	n1 := lib.TreeNode{1, &n2, &n3}
	depth := 0
	log.Println("平衡二叉树？", IsBalanceTree(&n1, &depth))
}
