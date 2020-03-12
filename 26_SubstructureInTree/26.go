package main

import "CodingInterview/getOffer-Go/lib"

func HasSubTree(pRoot1, pRoot2 *lib.TreeNode) bool {
	res := false
	if pRoot1 != nil && pRoot2 != nil {
		if pRoot1.Val == pRoot2.Val {
			res = tree1HaveTree2(pRoot1, pRoot2)
		}
		if !res {
			res = tree1HaveTree2(pRoot1.Left, pRoot2)
		}
		if !res {
			res = tree1HaveTree2(pRoot1.Right, pRoot2)
		}
	}
	return res
}

// 判断tree2是不是tree1的子树
func tree1HaveTree2(pRoot1, pRoot2 *lib.TreeNode) bool {
	if pRoot2 == nil {
		return true
	}
	if pRoot1 == nil {
		return false
	}
	if pRoot1.Val != pRoot2.Val {
		return false
	}
	return tree1HaveTree2(pRoot1.Left, pRoot2.Left) && tree1HaveTree2(pRoot1.Right, pRoot2.Right)
}
