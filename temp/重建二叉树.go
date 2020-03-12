package main

import "fmt"

/*
二叉树示例图.jpg
	前序遍历序列{1,2,4,7,3,5,6,8}
	中序遍历序列{4,7,2,1,5,3,8,6}

思路总结：
	先根据前序遍历序列的第一个数字创建根结点，
	接下来在中序遍历序列中找到根结点的位置，
	这样就能确定左、右子树结点的数量。
	在前序遍历和中序遍历的序列中划分了
	左、右子树结点的值之后，
	就可以递归地去分别构建它的左右子树。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(pre, in []int) *TreeNode {
	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: pre[0],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = buildTree(pre[1:idx+1], in[:idx])
	res.Right = buildTree(pre[idx+1:], in[idx+1:])

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

func main() {
	preorder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	inorder := []int{4, 7, 2, 1, 5, 3, 8, 6}
	res := buildTree(preorder, inorder)
	fmt.Print("binary tree:", res)
}
