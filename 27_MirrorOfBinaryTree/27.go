package main

import "CodingInterview/getOffer-Go/lib"

func GetMirror(pNode *lib.TreeNode) {
	if pNode == nil {
		return
	}
	if pNode.Left == nil && pNode.Right == nil {
		return
	}
	pTemp := pNode.Left
	pNode.Left = pNode.Right
	pNode.Right = pTemp

	if pNode.Left != nil {
		GetMirror(pNode.Left)
	}
	if pNode.Right != nil {
		GetMirror(pNode.Right)
	}
}
