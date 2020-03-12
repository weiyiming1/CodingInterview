package main

import "CodingInterview/getOffer-Go/lib"

func Convert(pRoot *lib.TreeNode) *lib.TreeNode {
	var pLastNodeInList *lib.TreeNode
	convertNode(pRoot, &pLastNodeInList)

	// pLastNodeInList 指向双向链表的尾节点
	// 题目要求返回头节点
	pHeadOfList := pLastNodeInList
	for pHeadOfList != nil && pHeadOfList.Left != nil {
		pHeadOfList = pHeadOfList.Left
	}

	return pHeadOfList
}

func convertNode(pNode *lib.TreeNode, pLastNodeInList **lib.TreeNode) {
	if pNode == nil {
		return
	}
	pCur := pNode
	if pCur.Left != nil {
		convertNode(pCur.Left, pLastNodeInList)
	}
	pCur.Left = *pLastNodeInList
	if *pLastNodeInList != nil {
		(*pLastNodeInList).Right = pCur
	}
	*pLastNodeInList = pCur
	if pCur.Right != nil {
		convertNode(pCur.Right, pLastNodeInList)
	}
}
