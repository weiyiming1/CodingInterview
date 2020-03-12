package main

import "CodingInterview/getOffer-Go/lib"

func Symmetrical(pRoot *lib.TreeNode) bool {
	return getSymmetrical(pRoot, pRoot)
}

func getSymmetrical(pRoot1, pRoot2 *lib.TreeNode) bool {
	if pRoot1 == nil && pRoot2 == nil {
		return true
	}
	if pRoot1 == nil || pRoot2 == nil {
		return false
	}
	if pRoot1.Val != pRoot2.Val {
		return false
	}
	return getSymmetrical(pRoot1.Left, pRoot2.Right) &&
		getSymmetrical(pRoot1.Right, pRoot2.Left)
}
