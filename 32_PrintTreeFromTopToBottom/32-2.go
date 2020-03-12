package main

import (
	"CodingInterview/getOffer-Go/lib"
	"fmt"
)

func PrintByLine(pRoot *lib.TreeNode) {
	if pRoot == nil {
		return
	}

	var queue []*lib.TreeNode
	queue = append(queue, pRoot)
	toBePrint := 1 // 当前层还没有打印的节点数
	nextLevel := 0 // 下一层的节点数
	for len(queue) > 0 {
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
			nextLevel++
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
			nextLevel++
		}
		fmt.Print(queue[0].Val)
		queue = queue[1:]
		toBePrint--
		if toBePrint == 0 {
			fmt.Print("\n")
			toBePrint = nextLevel
			nextLevel = 0
		}
	}
}
