package main

type Node struct {
	Val  int
	Next *Node
}

func FindKthToTail(node *Node, k int) *Node {
	if node == nil || k == 0 { //倒数计数从1开始
		return nil
	}

	pAhead := node
	for i := 0; i < k-1; i++ {
		if pAhead.Next != nil {
			pAhead = pAhead.Next
		} else {
			return nil
		}
	}
	pBhind := node
	for pAhead.Next != nil {
		pAhead = pAhead.Next
		pBhind = pBhind.Next
	}
	return pBhind
}
