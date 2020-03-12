package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用3个指针,逐个链接点进行翻转
func myReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { // 少于2个节点没有反转的必要
		return head
	}
	p := head
	q := head.Next
	head.Next = nil // 旧的头指针是新的尾指针，next需要指向nil
	for q != nil {
		r := q.Next // 保留下一个step要处理的指针
		q.Next = p
		p = q
		q = r
	}
	head = p // 最后q必然指向nil, 返回p作为新的头指针
	return head
}

func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	tmp := myReverseList(head)
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}
