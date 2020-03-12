package main

import (
	"fmt"
	"testing"
)

func TestFindFirstCommonNode(t *testing.T) {
	node7 := Node{7, nil}
	node6 := Node{6, &node7}
	node5 := Node{5, &node6}
	node4 := Node{4, &node5}
	node3 := Node{3, &node6}
	node2 := Node{2, &node3}
	node1 := Node{1, &node2}
	pCom := FindFirstCommonNode(&node1, &node4)
	fmt.Println(pCom)
}
