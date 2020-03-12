package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

//               max(maxDepth(9), maxDepth(20))+1
// max(maxDeth(nil), maxDeth(nil))+1            max(maxDeth(15), maxDeth(7))+1
//   1								max(maxDeth(nil), maxDeth(nil))+1         max(maxDeth(nil), maxDeth(nil))+1
//													1									1
