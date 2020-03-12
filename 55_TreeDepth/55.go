package main

import "CodingInterview/getOffer-Go/lib"

func TreeDepth(pRoot *lib.TreeNode) int {
	if pRoot == nil {
		return 0
	}

	left := TreeDepth(pRoot.Left)
	right := TreeDepth(pRoot.Right)

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
