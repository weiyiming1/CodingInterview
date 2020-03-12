package main

import (
	"CodingInterview/getOffer-Go/lib"
	"fmt"
)

func FindPath(pRoot *lib.TreeNode, target int) {
	if pRoot == nil {
		return
	}

	var path []int
	cur := 0
	getPath(pRoot, path, target, cur)

}

func getPath(pRoot *lib.TreeNode, path []int, target int, cur int) {
	cur += pRoot.Val
	path = append(path, pRoot.Val)
	isLeaf := false

	// 如果是叶节点，并且节点的值等于target的值，则打印这条路径
	if pRoot.Left == nil && pRoot.Right == nil {
		isLeaf = true
	}
	if cur == target && isLeaf {
		fmt.Println("找到路径:")
		for _, temp := range path {
			fmt.Println(temp)
		}
	}

	// 如果不是叶节点， 则遍历它的子节点
	if pRoot.Left != nil {
		getPath(pRoot.Left, path, target, cur)
	}
	if pRoot.Right != nil {
		getPath(pRoot.Right, path, target, cur)
	}

	// 返回父节点前， 在路径上删除当前节点
	path = path[:len(path)-1]
}
