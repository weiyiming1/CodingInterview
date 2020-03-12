package main

type Node struct {
	Val  int
	Next *Node
}

func DeleteNode(head, cur *Node) {
	if &head == nil || &cur == nil {
		return
	}

	if cur.Next != nil { // 非尾节点
		cur.Val = cur.Next.Val
		cur.Next = cur.Next.Next
	} else if head == cur { // 只有一个节点，删除头节点
		*head = Node{}
	} else { // 尾节点
		pNode := head
		for pNode.Next != cur {
			pNode = pNode.Next
		}
		pNode.Next = nil
	}
}
