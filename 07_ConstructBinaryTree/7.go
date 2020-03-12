package main

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func ReConstructTree(preOrder, inOrder []int) *BinaryTreeNode {
	if &preOrder == nil || &inOrder == nil || len(preOrder) <= 0 {
		return nil
	}

	res := &BinaryTreeNode{
		Val: preOrder[0],
	}

	idx := indexOf(res.Val, inOrder)

	res.Left = ReConstructTree(preOrder[1:idx+1], inOrder[:idx])
	res.Right = ReConstructTree(preOrder[idx+1:], inOrder[idx+1:])

	return res
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}
	return 0
}
