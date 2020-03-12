package main

type Node struct {
	Val  int
	Next *Node
}

func FindLoopEntry(pHead *Node) *Node {
	meetN := getMeetNode(pHead)
	if meetN == nil {
		return nil
	}
	//获取环的点数
	loopNum := 1
	pNode := meetN
	for pNode.Next != meetN {
		pNode = pNode.Next
		loopNum++
	}

	// 先移动pNode, 次数为loopNum
	pNode = pHead
	for i := 0; i < loopNum; i++ {
		pNode = pNode.Next
	}

	//移动pNode和pNode2
	pNode2 := pHead
	for pNode != pNode2 {
		pNode = pNode.Next
		pNode2 = pNode2.Next
	}
	return pNode
}

func getMeetNode(pHead *Node) *Node {
	// 空链表
	if pHead == nil {
		return nil
	}
	pSlow := pHead.Next
	// 只有一个节点，不可能有环
	if pSlow == nil {
		return nil
	}
	pFast := pSlow.Next
	for pSlow != nil && pFast != nil {
		if pFast == pSlow {
			return pFast
		}
		pSlow = pSlow.Next
		pFast = pFast.Next
		if pFast != nil {
			pFast = pFast.Next
		}
	}
	return nil
}
