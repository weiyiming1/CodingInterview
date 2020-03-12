package main

import (
	"CodingInterview/getOffer-Go/lib"
	"log"
	"testing"
)

func TestKthNodeInBST(t *testing.T) {
	n2 := lib.TreeNode{2, nil, nil}
	n4 := lib.TreeNode{4, nil, nil}
	n6 := lib.TreeNode{6, nil, nil}
	n8 := lib.TreeNode{8, nil, nil}
	n3 := lib.TreeNode{3, &n2, &n4}
	n7 := lib.TreeNode{7, &n6, &n8}
	n5 := lib.TreeNode{5, &n3, &n7}
	log.Printf("第%d大节点是%v：", 3, *KthNodeInBST(&n5, 3))
}
