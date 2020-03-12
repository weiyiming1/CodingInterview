package main

import (
	"CodingInterview/getOffer-Go/lib"
)

// BST binary search tree
func KthNodeInBST(pRoot *lib.TreeNode, k int) *lib.TreeNode {
	if pRoot == nil || k == 0 {
		return nil
	}
	var seq []lib.TreeNode
	inOrder(pRoot, &seq)
	for idx, val := range seq {
		if idx+1 == k {
			return &val
		}
	}
	return nil
}

func inOrder(pRoot *lib.TreeNode, seq *[]lib.TreeNode) {
	if pRoot == nil {
		return
	}
	if pRoot.Left != nil {
		inOrder(pRoot.Left, seq)
	}

	*seq = append(*seq, *pRoot)

	if pRoot != nil {
		inOrder(pRoot.Right, seq)
	}
}

/* 解法2：
func getKthNode(pRoot *lib.TreeNode, k *int)*lib.TreeNode{
	var target *lib.TreeNode
	if pRoot.Left != nil{
		target = getKthNode(pRoot.Left, k)
	}

	// 没有左子节点
	if target == nil{
		if *k==1{
			return pRoot
		}
		*k--
	}

	// 右子节点
	if target == nil && pRoot.Right != nil{
		target = getKthNode(pRoot.Right, k)
	}
	return target
}*/
