package main

type Node struct {
	Val  int
	Next *Node
}

func Merge(pHead1, pHead2 *Node) *Node {
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}
	var pMergedHead *Node
	if pHead1.Val < pHead2.Val {
		pMergedHead = pHead1
		pMergedHead.Next = Merge(pHead1.Next, pHead2)
	} else {
		pMergedHead = pHead2
		pMergedHead.Next = Merge(pHead1, pHead2.Next)
	}
	return pMergedHead
}
