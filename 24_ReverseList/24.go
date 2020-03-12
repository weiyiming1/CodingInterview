package main

type Node struct {
	Val  int
	Next *Node
}

func ReverseList(pHead *Node) *Node {
	pNode := pHead
	var pPre *Node
	var pReverseHead *Node
	for pNode != nil {
		pNext := pNode.Next
		if pNext == nil {
			pReverseHead = pNode
		}
		pNode.Next = pPre
		pPre = pNode
		pNode = pNext
	}
	return pReverseHead
}
