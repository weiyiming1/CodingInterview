package main

import (
	"CodingInterview/getOffer-Go/lib"
	"fmt"
)

func PrintFromTopToBottom(pRoot *lib.TreeNode) {
	if pRoot == nil {
		return
	}
	var queue []*lib.TreeNode
	queue = append(queue, pRoot)
	for len(queue) > 0 {
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}
		fmt.Println(*queue[0])
		queue = queue[1:]
	}
}
