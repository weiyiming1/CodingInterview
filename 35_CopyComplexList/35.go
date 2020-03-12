package main

import "fmt"

type ComplexListNode struct {
	Val     int
	Next    *ComplexListNode
	Sibling *ComplexListNode
}

func cloneNodes(pHead *ComplexListNode) {
	pNode := pHead
	for pNode != nil {
		pCloneNode := new(ComplexListNode)
		fmt.Println("-------")
		fmt.Println(pCloneNode)
		pCloneNode.Val = pNode.Val
		pCloneNode.Next = pNode.Next

		pNode.Next = pCloneNode
		pNode = pCloneNode.Next
	}
}

func connectSiblingNodes(pHead *ComplexListNode) {
	pNode := pHead
	for pNode != nil {
		pCloneNode := pNode.Next
		if pNode.Sibling != nil {
			pCloneNode.Sibling = pNode.Sibling.Next
		}
		pNode = pCloneNode.Next
	}
}

func reconnectNodes(pHead *ComplexListNode) *ComplexListNode {
	pNode := pHead
	var pClonedNode *ComplexListNode
	var pClonedHead *ComplexListNode
	if pNode != nil {
		pClonedHead = pNode.Next
		pClonedNode = pNode.Next
		pNode.Next = pClonedNode.Next
		pNode = pNode.Next
	}

	for pNode != nil {
		pClonedNode.Next = pNode.Next
		pClonedNode = pClonedNode.Next
		pNode.Next = pClonedNode.Next
		pNode = pNode.Next
	}

	return pClonedHead
}

func Clone(pHead *ComplexListNode) *ComplexListNode {
	cloneNodes(pHead)
	connectSiblingNodes(pHead)
	return reconnectNodes(pHead)
}
