package main

import "CodingInterview/getOffer-Go/lib"

func IsBalanceTree(pRoot *lib.TreeNode, pDepth *int) bool {
	if pRoot == nil {
		*pDepth = 0
		return true
	}

	var left, right int
	if IsBalanceTree(pRoot.Left, &left) && IsBalanceTree(pRoot.Right, &right) {
		dif := left - right
		if -1 <= dif && dif <= 1 {
			if left > right {
				*pDepth = 1 + left
			} else {
				*pDepth = 1 + right
			}
			return true
		}
	}
	return false
}
