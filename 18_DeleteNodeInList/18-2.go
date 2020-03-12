package main

// t: 所有元素都重复 bug unsolved !!!
func DeleteDuplicate(head *Node) {
	// head 辅助节点，next指向当前的链表的第一个节点
	if head.Next == nil { // 空链表为空
		return
	}
	var preNode *Node //用来记录重复节点之前的节点
	pNode := head.Next
	for pNode != nil {
		pNext := pNode.Next
		needDelete := false
		if pNext != nil && pNext.Val == pNode.Val {
			needDelete = true
		}
		if !needDelete {
			preNode = pNode
			pNode = pNode.Next // 更新当前节点
		} else {
			val := pNode.Val
			needDelNode := pNode
			for needDelNode != nil && needDelNode.Val == val {
				pNext = needDelNode.Next
				needDelNode = pNext
			}
			pNode = needDelNode
			if preNode == nil {
				head.Next = pNext
			} else {
				preNode.Next = pNext
			}
			pNode = pNext
		}
	}
}
