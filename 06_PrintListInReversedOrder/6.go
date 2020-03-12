package main

import "fmt"

// 递归
func PrintLinkReverse(Node *LinkNode) {
	if Node != nil {
		if Node.Next != nil {
			PrintLinkReverse(Node.Next)
		}
		fmt.Println(Node.Data)
	}
}

// 基于栈的循环
func PrintLinkReverseIterate(Node *LinkNode) {
	// todo
}
