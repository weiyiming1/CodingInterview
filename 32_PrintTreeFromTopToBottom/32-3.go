package main

import (
	"CodingInterview/getOffer-Go/lib"
	"fmt"
)

func ZPrint(pRoot *lib.TreeNode) {
	if pRoot == nil {
		return
	}

	var levels [2]lib.Stack // 使用2个栈结构
	cur := 0
	next := 1 // 使用cur和next来对栈进行索引

	levels[cur].Push(pRoot)
	for levels[cur].Size() > 0 || levels[next].Size() > 0 { // 栈还有元素

		pNode := levels[cur].Top().(*lib.TreeNode) // 入栈元素的父节点
		levels[cur].Pop()
		fmt.Print(pNode.Val, "\t")

		if cur == 0 { //左->右
			if pNode.Left != nil {
				levels[next].Push(pNode.Left)
			}
			if pNode.Right != nil {
				levels[next].Push(pNode.Right)
			}
		} else { // 右->左
			if pNode.Right != nil {
				levels[next].Push(pNode.Right)
			}
			if pNode.Left != nil {
				levels[next].Push(pNode.Left)
			}
		}

		if levels[cur].Empty() {
			fmt.Print("\n")
			cur = 1 - cur
			next = 1 - next
		}
	}
}
