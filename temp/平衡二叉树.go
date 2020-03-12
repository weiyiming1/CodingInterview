package main

// 平衡二叉树：
// 		左右子树的高度之差不大于 1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if isBalanced(root.Left) == false {
		return false
	}
	if isBalanced(root.Right) == false {
		return false
	}
	if abs(maxDepth(root.Left)-maxDepth(root.Right)) > 1 {
		return false
	}
	return true
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// --------------------------
/*
func main(){
	node4 := TreeNode{Val: 4, Left: nil, Right: nil}
	node7 := TreeNode{Val: 7, Left: nil, Right: nil}
	node5 := TreeNode{Val: 5, Left: &node4, Right: &node7}

	fmt.Println(isBalanced(node5))
}*/
